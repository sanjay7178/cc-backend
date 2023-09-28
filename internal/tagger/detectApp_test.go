// Copyright (C) 2022 NHR@FAU, University Erlangen-Nuremberg.
// All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package tagger

import (
	"testing"

	"github.com/ClusterCockpit/cc-backend/internal/repository"
	"github.com/ClusterCockpit/cc-backend/pkg/log"
)

func setup(tb testing.TB) *repository.JobRepository {
	tb.Helper()
	log.Init("warn", true)
	dbfile := "../repository/testdata/job.db"
	err := repository.MigrateDB("sqlite3", dbfile)
	noErr(tb, err)
	repository.Connect("sqlite3", dbfile)
	return repository.GetJobRepository()
}

func noErr(tb testing.TB, err error) {
	tb.Helper()

	if err != nil {
		tb.Fatal("Error is not nil:", err)
	}
}

func TestRegister(t *testing.T) {
	var tagger AppTagger

	err := tagger.Register()
	noErr(t, err)

	if len(tagger.apps) != 3 {
		t.Errorf("wrong summary for diagnostic \ngot: %d \nwant: 3", len(tagger.apps))
	}
}

func TestMatch(t *testing.T) {
	r := setup(t)

	job, err := r.FindById(5)
	noErr(t, err)

	var tagger AppTagger

	err = tagger.Register()
	noErr(t, err)

	tagger.Match(job)

	if !r.HasTag(5, "app", "vasp") {
		t.Errorf("missing tag vasp")
	}
}
