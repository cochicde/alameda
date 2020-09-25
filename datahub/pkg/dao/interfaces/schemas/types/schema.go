package types

import (
	"prophetstor.com/alameda/pkg/database/influxdb/schemas"
)

type SchemaDAO interface {
	CreateSchemas(*CreateSchemasRequest) error
	ListSchemas(*ListSchemasRequest) ([]*schemas.Schema, error)
	DeleteSchemas(*DeleteSchemasRequest) error
}

type CreateSchemasRequest struct {
	Schemas []*schemas.Schema
}

type ListSchemasRequest struct {
	SchemaMeta *schemas.SchemaMeta
}

type DeleteSchemasRequest struct {
	SchemaMeta *schemas.SchemaMeta
}
