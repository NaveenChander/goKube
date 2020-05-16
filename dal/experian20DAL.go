package dal

import (
	_ "github.com/denisenkom/go-mssqldb"
)

type IExperian interface{}

// ExperianSQLDAL ... ExperianSQLDAL - SQL Call for Experian
type ExperianSQLDAL struct{}

// ExperianTestDAL ... ExperianTestDAL - Mock Call for Experian
type ExperianTestDAL struct{}
