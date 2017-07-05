package config

import (
	"testing"
)

func TestDbConfig_Import(t *testing.T) {
	InitConsulClientViaEnv()

	db := new(DbConfig)
	err := db.Import("user")
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", *db)
}
