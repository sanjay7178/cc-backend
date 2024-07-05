// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
)

const createFileStashURL = `-- name: CreateFileStashURL :exec
INSERT INTO file_stash_url (url)
VALUES (?)
ON DUPLICATE KEY UPDATE url = VALUES(url)
`

// File Stash URL
func (q *Queries) CreateFileStashURL(ctx context.Context, url string) error {
	_, err := q.db.ExecContext(ctx, createFileStashURL, url)
	return err
}

const createInfluxDBConfiguration = `-- name: CreateInfluxDBConfiguration :exec
INSERT INTO influxdb_configurations (
    type, database_name, host, port, user, password, organization,
    ssl_enabled, batch_size, retry_interval, retry_exponential_base,
    max_retries, max_retry_time, meta_as_tags
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
ON DUPLICATE KEY UPDATE
    type = VALUES(type),
    database_name = VALUES(database_name),
    host = VALUES(host),
    port = VALUES(port),
    user = VALUES(user),
    password = VALUES(password),
    organization = VALUES(organization),
    ssl_enabled = VALUES(ssl_enabled),
    batch_size = VALUES(batch_size),
    retry_interval = VALUES(retry_interval),
    retry_exponential_base = VALUES(retry_exponential_base),
    max_retries = VALUES(max_retries),
    max_retry_time = VALUES(max_retry_time),
    meta_as_tags = VALUES(meta_as_tags)
`

type CreateInfluxDBConfigurationParams struct {
	Type                 string
	DatabaseName         string
	Host                 string
	Port                 int32
	User                 string
	Password             string
	Organization         string
	SslEnabled           bool
	BatchSize            int32
	RetryInterval        string
	RetryExponentialBase int32
	MaxRetries           int32
	MaxRetryTime         string
	MetaAsTags           sql.NullString
}

// InfluxDB Configurations
func (q *Queries) CreateInfluxDBConfiguration(ctx context.Context, arg CreateInfluxDBConfigurationParams) error {
	_, err := q.db.ExecContext(ctx, createInfluxDBConfiguration,
		arg.Type,
		arg.DatabaseName,
		arg.Host,
		arg.Port,
		arg.User,
		arg.Password,
		arg.Organization,
		arg.SslEnabled,
		arg.BatchSize,
		arg.RetryInterval,
		arg.RetryExponentialBase,
		arg.MaxRetries,
		arg.MaxRetryTime,
		arg.MetaAsTags,
	)
	return err
}

const createLVMConf = `-- name: CreateLVMConf :exec
INSERT INTO lvm_conf (machine_id, username, minAvailableSpaceGB, maxAvailableSpaceGB)
VALUES (?, ?, ?, ?)
`

type CreateLVMConfParams struct {
	MachineID           string
	Username            string
	Minavailablespacegb float64
	Maxavailablespacegb float64
}

// LVM Conf
func (q *Queries) CreateLVMConf(ctx context.Context, arg CreateLVMConfParams) error {
	_, err := q.db.ExecContext(ctx, createLVMConf,
		arg.MachineID,
		arg.Username,
		arg.Minavailablespacegb,
		arg.Maxavailablespacegb,
	)
	return err
}

const createLVStorageIssuer = `-- name: CreateLVStorageIssuer :exec
INSERT INTO lv_storage_issuer (machine_id, inc_buffer, dec_buffer, hostname, username, minAvailableSpaceGB, maxAvailableSpaceGB)
VALUES (?, ?, ?, ?, ?, ?, ?)
`

type CreateLVStorageIssuerParams struct {
	MachineID           string
	IncBuffer           sql.NullInt32
	DecBuffer           sql.NullInt32
	Hostname            string
	Username            string
	Minavailablespacegb float64
	Maxavailablespacegb float64
}

// LV Storage Issuer
func (q *Queries) CreateLVStorageIssuer(ctx context.Context, arg CreateLVStorageIssuerParams) error {
	_, err := q.db.ExecContext(ctx, createLVStorageIssuer,
		arg.MachineID,
		arg.IncBuffer,
		arg.DecBuffer,
		arg.Hostname,
		arg.Username,
		arg.Minavailablespacegb,
		arg.Maxavailablespacegb,
	)
	return err
}

const createLogicalVolume = `-- name: CreateLogicalVolume :exec
INSERT INTO logical_volumes (machine_id, lv_name, vg_name, lv_attr, lv_size)
VALUES (?, ?, ?, ?, ?)
`

type CreateLogicalVolumeParams struct {
	MachineID string
	LvName    string
	VgName    string
	LvAttr    string
	LvSize    string
}

// Logical Volumes
func (q *Queries) CreateLogicalVolume(ctx context.Context, arg CreateLogicalVolumeParams) error {
	_, err := q.db.ExecContext(ctx, createLogicalVolume,
		arg.MachineID,
		arg.LvName,
		arg.VgName,
		arg.LvAttr,
		arg.LvSize,
	)
	return err
}

const createMachine = `-- name: CreateMachine :exec
INSERT INTO machines (machine_id, hostname, os_version, ip_address)
VALUES (?, ?, ?, ?)
`

type CreateMachineParams struct {
	MachineID string
	Hostname  string
	OsVersion string
	IpAddress string
}

// Machines
func (q *Queries) CreateMachine(ctx context.Context, arg CreateMachineParams) error {
	_, err := q.db.ExecContext(ctx, createMachine,
		arg.MachineID,
		arg.Hostname,
		arg.OsVersion,
		arg.IpAddress,
	)
	return err
}

const createMachineConf = `-- name: CreateMachineConf :exec
INSERT INTO machine_conf (machine_id, hostname, username, passphrase, port_number, password, host_key, folder_path)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateMachineConfParams struct {
	MachineID  string
	Hostname   string
	Username   string
	Passphrase sql.NullString
	PortNumber int32
	Password   sql.NullString
	HostKey    sql.NullString
	FolderPath sql.NullString
}

// Machine Conf
func (q *Queries) CreateMachineConf(ctx context.Context, arg CreateMachineConfParams) error {
	_, err := q.db.ExecContext(ctx, createMachineConf,
		arg.MachineID,
		arg.Hostname,
		arg.Username,
		arg.Passphrase,
		arg.PortNumber,
		arg.Password,
		arg.HostKey,
		arg.FolderPath,
	)
	return err
}

const createNotification = `-- name: CreateNotification :exec
INSERT INTO notifications (message) VALUES (?)
`

// Notifications
func (q *Queries) CreateNotification(ctx context.Context, message string) error {
	_, err := q.db.ExecContext(ctx, createNotification, message)
	return err
}

const createPhysicalVolume = `-- name: CreatePhysicalVolume :exec
INSERT INTO physical_volumes (machine_id, pv_name, vg_name, pv_fmt, pv_attr, pv_size, pv_free)
VALUES (?, ?, ?, ?, ?, ?, ?)
`

type CreatePhysicalVolumeParams struct {
	MachineID string
	PvName    string
	VgName    string
	PvFmt     string
	PvAttr    string
	PvSize    string
	PvFree    string
}

// Physical Volumes
func (q *Queries) CreatePhysicalVolume(ctx context.Context, arg CreatePhysicalVolumeParams) error {
	_, err := q.db.ExecContext(ctx, createPhysicalVolume,
		arg.MachineID,
		arg.PvName,
		arg.VgName,
		arg.PvFmt,
		arg.PvAttr,
		arg.PvSize,
		arg.PvFree,
	)
	return err
}

const createRabbitMQConfig = `-- name: CreateRabbitMQConfig :exec
INSERT INTO rabbit_mq_config (conn_url, username, password)
VALUES (?, ?, ?)
ON DUPLICATE KEY UPDATE conn_url = VALUES(conn_url), username = VALUES(username), password = VALUES(password)
`

type CreateRabbitMQConfigParams struct {
	ConnUrl  string
	Username string
	Password string
}

// RabbitMQ Config
func (q *Queries) CreateRabbitMQConfig(ctx context.Context, arg CreateRabbitMQConfigParams) error {
	_, err := q.db.ExecContext(ctx, createRabbitMQConfig, arg.ConnUrl, arg.Username, arg.Password)
	return err
}

const createRealtimeLog = `-- name: CreateRealtimeLog :exec
INSERT INTO realtime_logs (log_message, machine_id) VALUES (?, ?)
`

type CreateRealtimeLogParams struct {
	LogMessage string
	MachineID  string
}

// Realtime Logs
func (q *Queries) CreateRealtimeLog(ctx context.Context, arg CreateRealtimeLogParams) error {
	_, err := q.db.ExecContext(ctx, createRealtimeLog, arg.LogMessage, arg.MachineID)
	return err
}

const createVolumeGroup = `-- name: CreateVolumeGroup :exec
INSERT INTO volume_groups (machine_id, vg_name, pv_count, lv_count, snap_count, vg_attr, vg_size, vg_free)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateVolumeGroupParams struct {
	MachineID string
	VgName    string
	PvCount   string
	LvCount   string
	SnapCount string
	VgAttr    string
	VgSize    string
	VgFree    string
}

// Volume Groups
func (q *Queries) CreateVolumeGroup(ctx context.Context, arg CreateVolumeGroupParams) error {
	_, err := q.db.ExecContext(ctx, createVolumeGroup,
		arg.MachineID,
		arg.VgName,
		arg.PvCount,
		arg.LvCount,
		arg.SnapCount,
		arg.VgAttr,
		arg.VgSize,
		arg.VgFree,
	)
	return err
}

const deleteFileStashURL = `-- name: DeleteFileStashURL :exec
DELETE FROM file_stash_url WHERE id = ?
`

func (q *Queries) DeleteFileStashURL(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteFileStashURL, id)
	return err
}

const deleteInfluxDBConfiguration = `-- name: DeleteInfluxDBConfiguration :exec
DELETE FROM influxdb_configurations WHERE id = ?
`

func (q *Queries) DeleteInfluxDBConfiguration(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteInfluxDBConfiguration, id)
	return err
}

const deleteLVMConf = `-- name: DeleteLVMConf :exec
DELETE FROM lvm_conf WHERE id = ?
`

func (q *Queries) DeleteLVMConf(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteLVMConf, id)
	return err
}

const deleteLVStorageIssuer = `-- name: DeleteLVStorageIssuer :exec
DELETE FROM lv_storage_issuer WHERE id = ?
`

func (q *Queries) DeleteLVStorageIssuer(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteLVStorageIssuer, id)
	return err
}

const deleteLogicalVolume = `-- name: DeleteLogicalVolume :exec
DELETE FROM logical_volumes
WHERE lv_id = ?
`

func (q *Queries) DeleteLogicalVolume(ctx context.Context, lvID int32) error {
	_, err := q.db.ExecContext(ctx, deleteLogicalVolume, lvID)
	return err
}

const deleteMachine = `-- name: DeleteMachine :exec
DELETE FROM machines
WHERE machine_id = ?
`

func (q *Queries) DeleteMachine(ctx context.Context, machineID string) error {
	_, err := q.db.ExecContext(ctx, deleteMachine, machineID)
	return err
}

const deleteMachineConf = `-- name: DeleteMachineConf :exec
DELETE FROM machine_conf WHERE id = ?
`

func (q *Queries) DeleteMachineConf(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteMachineConf, id)
	return err
}

const deleteNotification = `-- name: DeleteNotification :exec
DELETE FROM notifications WHERE id = ?
`

func (q *Queries) DeleteNotification(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteNotification, id)
	return err
}

const deletePhysicalVolume = `-- name: DeletePhysicalVolume :exec
DELETE FROM physical_volumes
WHERE pv_id = ?
`

func (q *Queries) DeletePhysicalVolume(ctx context.Context, pvID int32) error {
	_, err := q.db.ExecContext(ctx, deletePhysicalVolume, pvID)
	return err
}

const deleteRabbitMQConfig = `-- name: DeleteRabbitMQConfig :exec
DELETE FROM rabbit_mq_config WHERE single_row_enforcer = 1
`

func (q *Queries) DeleteRabbitMQConfig(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteRabbitMQConfig)
	return err
}

const deleteRealtimeLog = `-- name: DeleteRealtimeLog :exec
DELETE FROM realtime_logs WHERE id = ?
`

func (q *Queries) DeleteRealtimeLog(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteRealtimeLog, id)
	return err
}

const deleteVolumeGroup = `-- name: DeleteVolumeGroup :exec
DELETE FROM volume_groups
WHERE vg_id = ?
`

func (q *Queries) DeleteVolumeGroup(ctx context.Context, vgID int32) error {
	_, err := q.db.ExecContext(ctx, deleteVolumeGroup, vgID)
	return err
}

const getFileStashURL = `-- name: GetFileStashURL :one
SELECT id, url, created_at, single_row_enforcer FROM file_stash_url
LIMIT 1
`

func (q *Queries) GetFileStashURL(ctx context.Context) (FileStashUrl, error) {
	row := q.db.QueryRowContext(ctx, getFileStashURL)
	var i FileStashUrl
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.CreatedAt,
		&i.SingleRowEnforcer,
	)
	return i, err
}

const getInfluxDBConfiguration = `-- name: GetInfluxDBConfiguration :one
SELECT id, type, database_name, host, port, user, password, organization, ssl_enabled, batch_size, retry_interval, retry_exponential_base, max_retries, max_retry_time, meta_as_tags, single_row_enforcer FROM influxdb_configurations
LIMIT 1
`

func (q *Queries) GetInfluxDBConfiguration(ctx context.Context) (InfluxdbConfiguration, error) {
	row := q.db.QueryRowContext(ctx, getInfluxDBConfiguration)
	var i InfluxdbConfiguration
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.DatabaseName,
		&i.Host,
		&i.Port,
		&i.User,
		&i.Password,
		&i.Organization,
		&i.SslEnabled,
		&i.BatchSize,
		&i.RetryInterval,
		&i.RetryExponentialBase,
		&i.MaxRetries,
		&i.MaxRetryTime,
		&i.MetaAsTags,
		&i.SingleRowEnforcer,
	)
	return i, err
}

const getLVMConf = `-- name: GetLVMConf :one
SELECT id, machine_id, username, minavailablespacegb, maxavailablespacegb, created_at FROM lvm_conf
WHERE machine_id = ?
ORDER BY created_at DESC
LIMIT 1
`

func (q *Queries) GetLVMConf(ctx context.Context, machineID string) (LvmConf, error) {
	row := q.db.QueryRowContext(ctx, getLVMConf, machineID)
	var i LvmConf
	err := row.Scan(
		&i.ID,
		&i.MachineID,
		&i.Username,
		&i.Minavailablespacegb,
		&i.Maxavailablespacegb,
		&i.CreatedAt,
	)
	return i, err
}

const getLVStorageIssuers = `-- name: GetLVStorageIssuers :many
SELECT id, machine_id, inc_buffer, dec_buffer, hostname, username, minavailablespacegb, maxavailablespacegb FROM lv_storage_issuer
`

func (q *Queries) GetLVStorageIssuers(ctx context.Context) ([]LvStorageIssuer, error) {
	rows, err := q.db.QueryContext(ctx, getLVStorageIssuers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []LvStorageIssuer
	for rows.Next() {
		var i LvStorageIssuer
		if err := rows.Scan(
			&i.ID,
			&i.MachineID,
			&i.IncBuffer,
			&i.DecBuffer,
			&i.Hostname,
			&i.Username,
			&i.Minavailablespacegb,
			&i.Maxavailablespacegb,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLogicalVolumes = `-- name: GetLogicalVolumes :many
SELECT lv_id, machine_id, lv_name, vg_name, lv_attr, lv_size, created_at FROM logical_volumes
WHERE machine_id = ?
`

func (q *Queries) GetLogicalVolumes(ctx context.Context, machineID string) ([]LogicalVolume, error) {
	rows, err := q.db.QueryContext(ctx, getLogicalVolumes, machineID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []LogicalVolume
	for rows.Next() {
		var i LogicalVolume
		if err := rows.Scan(
			&i.LvID,
			&i.MachineID,
			&i.LvName,
			&i.VgName,
			&i.LvAttr,
			&i.LvSize,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMachine = `-- name: GetMachine :one
SELECT machine_id, hostname, os_version, ip_address, created_at FROM machines
WHERE machine_id = ?
`

func (q *Queries) GetMachine(ctx context.Context, machineID string) (Machine, error) {
	row := q.db.QueryRowContext(ctx, getMachine, machineID)
	var i Machine
	err := row.Scan(
		&i.MachineID,
		&i.Hostname,
		&i.OsVersion,
		&i.IpAddress,
		&i.CreatedAt,
	)
	return i, err
}

const getMachineConf = `-- name: GetMachineConf :one
SELECT id, machine_id, hostname, username, passphrase, port_number, password, host_key, folder_path FROM machine_conf
WHERE machine_id = ?
`

func (q *Queries) GetMachineConf(ctx context.Context, machineID string) (MachineConf, error) {
	row := q.db.QueryRowContext(ctx, getMachineConf, machineID)
	var i MachineConf
	err := row.Scan(
		&i.ID,
		&i.MachineID,
		&i.Hostname,
		&i.Username,
		&i.Passphrase,
		&i.PortNumber,
		&i.Password,
		&i.HostKey,
		&i.FolderPath,
	)
	return i, err
}

const getNotifications = `-- name: GetNotifications :many
SELECT id, message, created_at FROM notifications
ORDER BY created_at DESC
LIMIT ?
`

func (q *Queries) GetNotifications(ctx context.Context, limit int32) ([]Notification, error) {
	rows, err := q.db.QueryContext(ctx, getNotifications, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Notification
	for rows.Next() {
		var i Notification
		if err := rows.Scan(&i.ID, &i.Message, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPhysicalVolumes = `-- name: GetPhysicalVolumes :many
SELECT pv_id, machine_id, pv_name, vg_name, pv_fmt, pv_attr, pv_size, pv_free, created_at FROM physical_volumes
WHERE machine_id = ?
`

func (q *Queries) GetPhysicalVolumes(ctx context.Context, machineID string) ([]PhysicalVolume, error) {
	rows, err := q.db.QueryContext(ctx, getPhysicalVolumes, machineID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PhysicalVolume
	for rows.Next() {
		var i PhysicalVolume
		if err := rows.Scan(
			&i.PvID,
			&i.MachineID,
			&i.PvName,
			&i.VgName,
			&i.PvFmt,
			&i.PvAttr,
			&i.PvSize,
			&i.PvFree,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRabbitMQConfig = `-- name: GetRabbitMQConfig :one
SELECT conn_url, username, password, created_at, single_row_enforcer FROM rabbit_mq_config
LIMIT 1
`

func (q *Queries) GetRabbitMQConfig(ctx context.Context) (RabbitMqConfig, error) {
	row := q.db.QueryRowContext(ctx, getRabbitMQConfig)
	var i RabbitMqConfig
	err := row.Scan(
		&i.ConnUrl,
		&i.Username,
		&i.Password,
		&i.CreatedAt,
		&i.SingleRowEnforcer,
	)
	return i, err
}

const getRealtimeLogs = `-- name: GetRealtimeLogs :many
SELECT id, log_message, machine_id, created_at FROM realtime_logs
WHERE machine_id = ?
ORDER BY created_at DESC
LIMIT ?
`

type GetRealtimeLogsParams struct {
	MachineID string
	Limit     int32
}

func (q *Queries) GetRealtimeLogs(ctx context.Context, arg GetRealtimeLogsParams) ([]RealtimeLog, error) {
	rows, err := q.db.QueryContext(ctx, getRealtimeLogs, arg.MachineID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RealtimeLog
	for rows.Next() {
		var i RealtimeLog
		if err := rows.Scan(
			&i.ID,
			&i.LogMessage,
			&i.MachineID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getVolumeGroups = `-- name: GetVolumeGroups :many
SELECT vg_id, machine_id, vg_name, pv_count, lv_count, snap_count, vg_attr, vg_size, vg_free, created_at FROM volume_groups
WHERE machine_id = ?
`

func (q *Queries) GetVolumeGroups(ctx context.Context, machineID string) ([]VolumeGroup, error) {
	rows, err := q.db.QueryContext(ctx, getVolumeGroups, machineID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []VolumeGroup
	for rows.Next() {
		var i VolumeGroup
		if err := rows.Scan(
			&i.VgID,
			&i.MachineID,
			&i.VgName,
			&i.PvCount,
			&i.LvCount,
			&i.SnapCount,
			&i.VgAttr,
			&i.VgSize,
			&i.VgFree,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateFileStashURL = `-- name: UpdateFileStashURL :exec
UPDATE file_stash_url
SET url = ?
WHERE single_row_enforcer = 1
`

func (q *Queries) UpdateFileStashURL(ctx context.Context, url string) error {
	_, err := q.db.ExecContext(ctx, updateFileStashURL, url)
	return err
}

const updateInfluxDBConfiguration = `-- name: UpdateInfluxDBConfiguration :exec
UPDATE influxdb_configurations
SET type = ?, database_name = ?, host = ?, port = ?, user = ?, password = ?,
    organization = ?, ssl_enabled = ?, batch_size = ?, retry_interval = ?,
    retry_exponential_base = ?, max_retries = ?, max_retry_time = ?, meta_as_tags = ?
WHERE single_row_enforcer = 1
`

type UpdateInfluxDBConfigurationParams struct {
	Type                 string
	DatabaseName         string
	Host                 string
	Port                 int32
	User                 string
	Password             string
	Organization         string
	SslEnabled           bool
	BatchSize            int32
	RetryInterval        string
	RetryExponentialBase int32
	MaxRetries           int32
	MaxRetryTime         string
	MetaAsTags           sql.NullString
}

func (q *Queries) UpdateInfluxDBConfiguration(ctx context.Context, arg UpdateInfluxDBConfigurationParams) error {
	_, err := q.db.ExecContext(ctx, updateInfluxDBConfiguration,
		arg.Type,
		arg.DatabaseName,
		arg.Host,
		arg.Port,
		arg.User,
		arg.Password,
		arg.Organization,
		arg.SslEnabled,
		arg.BatchSize,
		arg.RetryInterval,
		arg.RetryExponentialBase,
		arg.MaxRetries,
		arg.MaxRetryTime,
		arg.MetaAsTags,
	)
	return err
}

const updateLVMConf = `-- name: UpdateLVMConf :exec
UPDATE lvm_conf
SET username = ?, minAvailableSpaceGB = ?, maxAvailableSpaceGB = ?
WHERE id = ?
`

type UpdateLVMConfParams struct {
	Username            string
	Minavailablespacegb float64
	Maxavailablespacegb float64
	ID                  int32
}

func (q *Queries) UpdateLVMConf(ctx context.Context, arg UpdateLVMConfParams) error {
	_, err := q.db.ExecContext(ctx, updateLVMConf,
		arg.Username,
		arg.Minavailablespacegb,
		arg.Maxavailablespacegb,
		arg.ID,
	)
	return err
}

const updateLVStorageIssuer = `-- name: UpdateLVStorageIssuer :exec
UPDATE lv_storage_issuer
SET inc_buffer = ?, dec_buffer = ?, hostname = ?, username = ?, minAvailableSpaceGB = ?, maxAvailableSpaceGB = ?
WHERE id = ?
`

type UpdateLVStorageIssuerParams struct {
	IncBuffer           sql.NullInt32
	DecBuffer           sql.NullInt32
	Hostname            string
	Username            string
	Minavailablespacegb float64
	Maxavailablespacegb float64
	ID                  int32
}

func (q *Queries) UpdateLVStorageIssuer(ctx context.Context, arg UpdateLVStorageIssuerParams) error {
	_, err := q.db.ExecContext(ctx, updateLVStorageIssuer,
		arg.IncBuffer,
		arg.DecBuffer,
		arg.Hostname,
		arg.Username,
		arg.Minavailablespacegb,
		arg.Maxavailablespacegb,
		arg.ID,
	)
	return err
}

const updateLogicalVolume = `-- name: UpdateLogicalVolume :exec
UPDATE logical_volumes
SET lv_name = ?, vg_name = ?, lv_attr = ?, lv_size = ?
WHERE lv_id = ?
`

type UpdateLogicalVolumeParams struct {
	LvName string
	VgName string
	LvAttr string
	LvSize string
	LvID   int32
}

func (q *Queries) UpdateLogicalVolume(ctx context.Context, arg UpdateLogicalVolumeParams) error {
	_, err := q.db.ExecContext(ctx, updateLogicalVolume,
		arg.LvName,
		arg.VgName,
		arg.LvAttr,
		arg.LvSize,
		arg.LvID,
	)
	return err
}

const updateMachine = `-- name: UpdateMachine :exec
UPDATE machines
SET hostname = ?, os_version = ?, ip_address = ?
WHERE machine_id = ?
`

type UpdateMachineParams struct {
	Hostname  string
	OsVersion string
	IpAddress string
	MachineID string
}

func (q *Queries) UpdateMachine(ctx context.Context, arg UpdateMachineParams) error {
	_, err := q.db.ExecContext(ctx, updateMachine,
		arg.Hostname,
		arg.OsVersion,
		arg.IpAddress,
		arg.MachineID,
	)
	return err
}

const updateMachineConf = `-- name: UpdateMachineConf :exec
UPDATE machine_conf
SET hostname = ?, username = ?, passphrase = ?, port_number = ?, password = ?, host_key = ?, folder_path = ?
WHERE id = ?
`

type UpdateMachineConfParams struct {
	Hostname   string
	Username   string
	Passphrase sql.NullString
	PortNumber int32
	Password   sql.NullString
	HostKey    sql.NullString
	FolderPath sql.NullString
	ID         int32
}

func (q *Queries) UpdateMachineConf(ctx context.Context, arg UpdateMachineConfParams) error {
	_, err := q.db.ExecContext(ctx, updateMachineConf,
		arg.Hostname,
		arg.Username,
		arg.Passphrase,
		arg.PortNumber,
		arg.Password,
		arg.HostKey,
		arg.FolderPath,
		arg.ID,
	)
	return err
}

const updatePhysicalVolume = `-- name: UpdatePhysicalVolume :exec
UPDATE physical_volumes
SET pv_name = ?, vg_name = ?, pv_fmt = ?, pv_attr = ?, pv_size = ?, pv_free = ?
WHERE pv_id = ?
`

type UpdatePhysicalVolumeParams struct {
	PvName string
	VgName string
	PvFmt  string
	PvAttr string
	PvSize string
	PvFree string
	PvID   int32
}

func (q *Queries) UpdatePhysicalVolume(ctx context.Context, arg UpdatePhysicalVolumeParams) error {
	_, err := q.db.ExecContext(ctx, updatePhysicalVolume,
		arg.PvName,
		arg.VgName,
		arg.PvFmt,
		arg.PvAttr,
		arg.PvSize,
		arg.PvFree,
		arg.PvID,
	)
	return err
}

const updateRabbitMQConfig = `-- name: UpdateRabbitMQConfig :exec
UPDATE rabbit_mq_config
SET conn_url = ?, username = ?, password = ?
WHERE single_row_enforcer = 1
`

type UpdateRabbitMQConfigParams struct {
	ConnUrl  string
	Username string
	Password string
}

func (q *Queries) UpdateRabbitMQConfig(ctx context.Context, arg UpdateRabbitMQConfigParams) error {
	_, err := q.db.ExecContext(ctx, updateRabbitMQConfig, arg.ConnUrl, arg.Username, arg.Password)
	return err
}

const updateVolumeGroup = `-- name: UpdateVolumeGroup :exec
UPDATE volume_groups
SET vg_name = ?, pv_count = ?, lv_count = ?, snap_count = ?, vg_attr = ?, vg_size = ?, vg_free = ?
WHERE vg_id = ?
`

type UpdateVolumeGroupParams struct {
	VgName    string
	PvCount   string
	LvCount   string
	SnapCount string
	VgAttr    string
	VgSize    string
	VgFree    string
	VgID      int32
}

func (q *Queries) UpdateVolumeGroup(ctx context.Context, arg UpdateVolumeGroupParams) error {
	_, err := q.db.ExecContext(ctx, updateVolumeGroup,
		arg.VgName,
		arg.PvCount,
		arg.LvCount,
		arg.SnapCount,
		arg.VgAttr,
		arg.VgSize,
		arg.VgFree,
		arg.VgID,
	)
	return err
}
