package repository

import (
	"context"
	"reflect"

	"main/services/profile-service/internal/values"

	"github.com/uptrace/bun"
)

type CRUD interface {
	GetByID(ctx context.Context, model any, relations ...string) error
	Create(ctx context.Context, model any) error
	Update(ctx context.Context, model any) error
	Delete(ctx context.Context, model any) error
}

type crud struct {
	idb bun.IDB
}

func NewCRUD(idb bun.IDB) CRUD {
	return &crud{
		idb: idb,
	}
}

func (c *crud) isNotAPointer(data any) error {
	if reflect.ValueOf(data).Kind() != reflect.Pointer {
		return values.ErrNotAPointer
	}
	return nil
}

func (c *crud) GetByID(ctx context.Context, model any, relations ...string) error {
	if err := c.isNotAPointer(model); err != nil {
		return err
	}
	q := c.idb.NewSelect().Model(model).WherePK()
	for i := range relations {
		q.Relation(relations[i])
	}
	return q.Scan(ctx)
}

func (c *crud) Create(ctx context.Context, model any) error {
	if err := c.isNotAPointer(model); err != nil {
		return err
	}
	_, err := c.idb.NewInsert().Model(model).Exec(ctx)
	return err
}

func (c *crud) Update(ctx context.Context, model any) error {
	if err := c.isNotAPointer(model); err != nil {
		return err
	}
	_, err := c.idb.NewUpdate().Model(model).WherePK().Exec(ctx)
	return err
}

func (c *crud) Delete(ctx context.Context, model any) error {
	if err := c.isNotAPointer(model); err != nil {
		return err
	}
	_, err := c.idb.NewDelete().Model(model).Exec(ctx)
	return err
}
