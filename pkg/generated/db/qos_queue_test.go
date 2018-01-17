package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
)

func TestQosQueue(t *testing.T) {
	t.Parallel()
	db := testDB
	common.UseTable(db, "metadata")
	common.UseTable(db, "qos_queue")
	defer func() {
		common.ClearTable(db, "qos_queue")
		common.ClearTable(db, "metadata")
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeQosQueue()
	model.UUID = "qos_queue_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "qos_queue_dummy"}
	model.Perms2.Owner = "admin"

	err := common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateQosQueue(tx, model)
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListQosQueue(tx, &common.ListSpec{Limit: 1})
		if err != nil {
			return err
		}
		if len(models) != 1 {
			return fmt.Errorf("expected one element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteQosQueue(tx, model.UUID,
			common.NewAuthContext("default", "demo", "demo", []string{}),
		)
	})
	if err == nil {
		t.Fatal("auth failed")
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteQosQueue(tx, model.UUID, nil)
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListQosQueue(tx, &common.ListSpec{Limit: 1})
		if err != nil {
			return err
		}
		if len(models) != 0 {
			return fmt.Errorf("expected no element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}
	return
}
