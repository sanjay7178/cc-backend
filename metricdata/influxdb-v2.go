package metricdata

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
	"crypto/tls"
	"encoding/json"

	"github.com/ClusterCockpit/cc-backend/config"
	"github.com/ClusterCockpit/cc-backend/schema"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	influxdb2Api "github.com/influxdata/influxdb-client-go/v2/api"
)

type InfluxDBv2DataRepositoryConfig struct {
	Url   string `json:"url"`
	Token string `json:"token"`
	Bucket string `json:"bucket"`
	Org string `json:"org"`
	SkipTls bool `json:"skiptls"`
}

type InfluxDBv2DataRepository struct {
	client              influxdb2.Client
	queryClient         influxdb2Api.QueryAPI
	bucket, measurement string
}

func (idb *InfluxDBv2DataRepository) Init(rawConfig json.RawMessage) error {
	var config InfluxDBv2DataRepositoryConfig
	if err := json.Unmarshal(rawConfig, &config); err != nil {
		return err
	}

	idb.client 			= influxdb2.NewClientWithOptions(config.Url, config.Token, influxdb2.DefaultOptions().SetTLSConfig(&tls.Config {InsecureSkipVerify: config.SkipTls,} ))
	idb.queryClient = idb.client.QueryAPI(config.Org)
	idb.bucket      = config.Bucket

	return nil
}

func (idb *InfluxDBv2DataRepository) formatTime(t time.Time) string {
	return t.Format(time.RFC3339) // Like “2006-01-02T15:04:05Z07:00”
}

func (idb *InfluxDBv2DataRepository) epochToTime(epoch int64) time.Time {
	return time.Unix(epoch, 0)
}

func (idb *InfluxDBv2DataRepository) LoadData(job *schema.Job, metrics []string, scopes []schema.MetricScope, ctx context.Context) (schema.JobData, error) {

	// DEBUG
	// log.Println("<< Requested Metrics >> ")
	// log.Println(metrics)
  // log.Println("<< Requested Scope >> ")
	// log.Println(scopes)

	measurementsConds := make([]string, 0, len(metrics))
	for _, m := range metrics {
		measurementsConds = append(measurementsConds, fmt.Sprintf(`r["_measurement"] == "%s"`, m))
	}
	measurementsCond := strings.Join(measurementsConds, " or ")

	hostsConds := make([]string, 0, len(job.Resources))
	for _, h := range job.Resources {
		if h.HWThreads != nil || h.Accelerators != nil {
			// TODO
			return nil, errors.New("the InfluxDB metric data repository does not yet support HWThreads or Accelerators")
		}

		hostsConds = append(hostsConds, fmt.Sprintf(`r["hostname"] == "%s"`, h.Hostname))
	}
	hostsCond := strings.Join(hostsConds, " or ")

	jobData := make(schema.JobData) // Empty Schema: map[<string>FIELD]map[<MetricScope>SCOPE]<*JobMetric>METRIC
	// Requested Scopes
	for _, scope := range scopes {

			// Query Influxdb
			query := ""

			switch scope {
					case "node":
							// Get Finest Granularity, Groupy By Measurement and Hostname (== Metric / Node), Calculate Mean, Set NULL to 0.0
							query = fmt.Sprintf(`
								from(bucket: "%s")
								|> range(start: %s, stop: %s)
								|> filter(fn: (r) => %s )
								|> filter(fn: (r) => %s )
								|> drop(columns: ["_start", "_stop"])
								|> group(columns: ["hostname", "_measurement"])
		            |> aggregateWindow(every: 60s, fn: mean)
								|> map(fn: (r) => (if exists r._value then {r with _value: r._value} else {r with _value: 0.0}))`,
								idb.bucket,
								idb.formatTime(job.StartTime), idb.formatTime(idb.epochToTime(job.StartTimeUnix + int64(job.Duration) + int64(1) )),
								measurementsCond, hostsCond)
					default:
							log.Println("Note: Other scope than 'node' requested, but not yet supported: Will return 'node' scope. ")
							query = fmt.Sprintf(`
								from(bucket: "%s")
								|> range(start: %s, stop: %s)
								|> filter(fn: (r) => %s )
								|> filter(fn: (r) => %s )
								|> drop(columns: ["_start", "_stop"])
								|> group(columns: ["hostname", "_measurement"])
								|> aggregateWindow(every: 60s, fn: mean)
								|> map(fn: (r) => (if exists r._value then {r with _value: r._value} else {r with _value: 0.0}))`,
								idb.bucket,
								idb.formatTime(job.StartTime), idb.formatTime(idb.epochToTime(job.StartTimeUnix + int64(job.Duration) + int64(1) )),
								measurementsCond, hostsCond)
							// return nil, errors.New("the InfluxDB metric data repository does not yet support other scopes than 'node'")
			}

			rows, err := idb.queryClient.Query(ctx, query)
			if err != nil {
				return nil, err
			}

			// Init Metrics
			for _, metric := range metrics {
					jobMetric, ok := jobData[metric]
					if !ok {
							mc 		:= config.GetMetricConfig(job.Cluster, metric)
							jobMetric = map[schema.MetricScope]*schema.JobMetric{
									scope: { // uses scope var from above!
											Unit:     mc.Unit,
											Scope:    scope,
											Timestep: mc.Timestep,
											Series:   make([]schema.Series, 0, len(job.Resources)),
											StatisticsSeries: nil, // Should be: &schema.StatsSeries{},
									},
							}
					}
					jobData[metric] = jobMetric
			}

			// Process Result: Time-Data
			field, host, hostSeries := "", "", schema.Series{}
			for rows.Next() {
				  row := rows.Record()
					if ( host == "" || host != row.ValueByKey("hostname").(string) || rows.TableChanged() ) {
					 		if ( host != "" ) {
									// Append Series before reset
					 		  	jobData[field][scope].Series = append(jobData[field][scope].Series, hostSeries)
					 		}
					 		field, host = row.Measurement(), row.ValueByKey("hostname").(string)
					 		hostSeries  = schema.Series{
					 				Hostname:   host,
					 				Statistics: nil,
					 				Data:       make([]schema.Float, 0),
					 		}
					}
					val := row.Value().(float64)
					hostSeries.Data = append(hostSeries.Data, schema.Float(val))
			}
			// Append last Series
		  jobData[field][scope].Series = append(jobData[field][scope].Series, hostSeries)
	}

	// Get Stats
	stats, err := idb.LoadStats(job, metrics, ctx)
	if err != nil {
		return nil, err
	}

	for _, scope := range scopes {
		for metric, nodes := range stats {
			// log.Println(fmt.Sprintf("<< Add Stats for : Field %s >>", metric))
			for node, stats := range nodes {
				// log.Println(fmt.Sprintf("<< Add Stats for : Host %s : Min %.2f, Max %.2f, Avg %.2f >>", node, stats.Min, stats.Max, stats.Avg ))
				for index, _ := range jobData[metric][scope].Series {
					// log.Println(fmt.Sprintf("<< Try to add Stats to Series in Position %d >>", index))
					if jobData[metric][scope].Series[index].Hostname == node {
						// log.Println(fmt.Sprintf("<< Match for Series in Position %d : Host %s >>", index, jobData[metric][scope].Series[index].Hostname))
						jobData[metric][scope].Series[index].Statistics = &schema.MetricStatistics{Avg: stats.Avg, Min: stats.Min, Max: stats.Max}
						// log.Println(fmt.Sprintf("<< Result Inner: Min %.2f, Max %.2f, Avg %.2f >>", jobData[metric][scope].Series[index].Statistics.Min, jobData[metric][scope].Series[index].Statistics.Max, jobData[metric][scope].Series[index].Statistics.Avg))
					}
				}
			}
		}
	}

	// DEBUG:
	// for _, scope := range scopes {
	// 		for _, met := range metrics {
	// 		   for _, series := range jobData[met][scope].Series {
	// 		   log.Println(fmt.Sprintf("<< Result: %d data points for metric %s on %s with scope %s, Stats: Min %.2f, Max %.2f, Avg %.2f >>",
	// 				 	len(series.Data), met, series.Hostname, scope,
	// 					series.Statistics.Min, series.Statistics.Max, series.Statistics.Avg))
	// 		   }
	// 		}
	// }

	return jobData, nil
}

func (idb *InfluxDBv2DataRepository) LoadStats(job *schema.Job, metrics []string, ctx context.Context) (map[string]map[string]schema.MetricStatistics, error) {

	stats := map[string]map[string]schema.MetricStatistics{}

	hostsConds := make([]string, 0, len(job.Resources))
	for _, h := range job.Resources {
			if h.HWThreads != nil || h.Accelerators != nil {
				// TODO
				return nil, errors.New("the InfluxDB metric data repository does not yet support HWThreads or Accelerators")
			}
			hostsConds = append(hostsConds, fmt.Sprintf(`r["hostname"] == "%s"`, h.Hostname))
	}
	hostsCond := strings.Join(hostsConds, " or ")

	for _, metric := range metrics {
			query := fmt.Sprintf(`
				  data = from(bucket: "%s")
				  |> range(start: %s, stop: %s)
				  |> filter(fn: (r) => r._measurement == "%s" and r._field == "value" and (%s))
				  union(tables: [data |> mean(column: "_value") |> set(key: "_field", value: "avg"),
				                 data |>  min(column: "_value") |> set(key: "_field", value: "min"),
				                 data |>  max(column: "_value") |> set(key: "_field", value: "max")])
				  |> pivot(rowKey: ["hostname"], columnKey: ["_field"], valueColumn: "_value")
				  |> group()`,
					idb.bucket,
					idb.formatTime(job.StartTime), idb.formatTime(idb.epochToTime(job.StartTimeUnix + int64(job.Duration) + int64(1) )),
					metric, hostsCond)

			rows, err := idb.queryClient.Query(ctx, query)
			if err != nil {
				return nil, err
			}

			nodes := map[string]schema.MetricStatistics{}
			for rows.Next() {
					row := rows.Record()
					host := row.ValueByKey("hostname").(string)
					avg, min, max := row.ValueByKey("avg").(float64),
						row.ValueByKey("min").(float64),
						row.ValueByKey("max").(float64)

					nodes[host] = schema.MetricStatistics{
						Avg: avg,
						Min: min,
						Max: max,
					}
			}
			stats[metric] = nodes
	}

	return stats, nil
}

func (idb *InfluxDBv2DataRepository) LoadNodeData(cluster, partition string, metrics, nodes []string, scopes []schema.MetricScope, from, to time.Time, ctx context.Context) (map[string]map[string][]*schema.JobMetric, error) {
	// TODO : Implement to be used in Analysis- und System/Node-View
	log.Println(fmt.Sprintf("LoadNodeData unimplemented for InfluxDBv2DataRepository, Args: cluster %s, partition %s, metrics %v, nodes %v, scopes %v", cluster, partition, metrics, nodes, scopes))

	return nil, errors.New("unimplemented for InfluxDBv2DataRepository")
}
