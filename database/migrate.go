package database

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/kmou424/resonance-dataserver/database/model"
	"github.com/kmou424/resonance-dataserver/internal/types"
)

func migrate() {
	migrator := Conn.Migrator()

	targets := []types.Pair[any, string]{
		{&model.GoodsMapper{}, model.GoodsMapperTable},
		{&model.AuthKey{}, model.AuthKeyTable},
	}

	for _, obj := range targets {
		if !migrator.HasTable(obj.First) && !migrator.HasTable(obj.Second) {
			err := migrator.CreateTable(obj.First)
			if err != nil {
				log.Fatal(fmt.Sprintf("create table [%s] failed", obj.Second))
			}
			if !migrator.HasTable(obj.Second) {
				err = migrator.RenameTable(obj.First, obj.Second)
				log.Fatal(fmt.Sprintf("rename table [%s] failed", obj.Second))
			}
		}
	}
}
