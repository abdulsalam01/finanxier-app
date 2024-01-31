package feature

import "github.com/api-sekejap/internal/repository/base"

const (
	tableDefinition = "features"

	// Select statements.
	querySelect     = `SELECT id, name, description,` + base.GetBaseAttrQuery + `FROM ` + tableDefinition
	querySelectByID = querySelect + ` WHERE id = ?`

	// Insert statements.
	queryInsert = `INSERT INTO ` + tableDefinition +
		`(name, description, is_active, created_by, updated_by) ` +
		`VALUES(?, ?, ?, ?, ?, ?, ?, ?)`
)
