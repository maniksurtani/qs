// Licensed under the Apache License, Version 2.0
// Details: https://raw.githubusercontent.com/maniksurtani/quotaservice/master/LICENSE

package quotaservice

import (
	"testing"
	"time"

	"github.com/maniksurtani/quotaservice/config"
	"github.com/maniksurtani/quotaservice/test/helpers"
)

func TestWithNoRpcs(t *testing.T) {
	helpers.ExpectingPanic(t, func() {
		New(&MockBucketFactory{}, &config.MemoryConfigPersister{})
	})
}

func TestValidServer(t *testing.T) {
	s := New(&MockBucketFactory{}, config.NewMemoryConfigPersister(), &MockEndpoint{})
	s.Start()
	defer s.Stop()
}

func TestUpdateConfig(t *testing.T) {
	p := config.NewMemoryConfigPersister()
	s := New(&MockBucketFactory{}, p, &MockEndpoint{}).(*server)

	originalConfig := config.NewDefaultServiceConfig()
	originalConfig.Version = 2
	originalConfig.Date = time.Now().Unix() - 10
	marshalledConfig, err := config.Marshal(originalConfig)

	if err != nil {
		t.Fatal("Error when updating config", err)
	}

	p.PersistAndNotify(marshalledConfig)

	s.Start()
	defer s.Stop()

	newConfig := config.NewDefaultServiceConfig()

	if err := s.UpdateConfig(newConfig, "test"); err != nil {
		t.Fatal("Error when updating config", err)
	}

	start := time.Now()

	for s.Configs().Version == originalConfig.Version {
		if time.Since(start) > time.Second {
			t.Fatal("Timeout waiting for config to change!")
		}

		time.Sleep(time.Millisecond * 5)
	}

	cfg := s.Configs()

	if cfg.User != "test" {
		t.Errorf("User %+v does not match passed in user \"test\"", s.cfgs.User)
	}

	if cfg.Version != 3 {
		t.Errorf("Version %+v does not match current version: 3", s.cfgs.Version)
	}

	if cfg.Date <= originalConfig.Date {
		t.Errorf("Date %+v was not updated from %+v", s.cfgs.Date, originalConfig.Date)
	}
}
