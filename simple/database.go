package simple

type Database struct {
	Name string
}
type DatabasePostgreSql Database
type DatabaseMongoDb Database

func NewDatabaseMongoDb() *DatabaseMongoDb {
	return (*DatabaseMongoDb)(&Database{Name: "MongoDb"})
}

func NewDatabasePostgreSql() *DatabasePostgreSql {
	return (*DatabasePostgreSql)(&Database{Name: "postgresql"})
}

type DatabaseRepository struct {
	DatabasePostgreSql *DatabasePostgreSql
	DatabaseMongoDb    *DatabaseMongoDb
}

func NewDatabaseRepository(postgreSql *DatabasePostgreSql, mongoDb *DatabaseMongoDb) *DatabaseRepository {
	return &DatabaseRepository{DatabasePostgreSql: postgreSql, DatabaseMongoDb: mongoDb}
}
