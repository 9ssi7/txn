module github.com/9ssi7/txn/examples/basic-gorm

replace github.com/9ssi7/txn => ../..

replace github.com/9ssi7/txngorm => ../../txngorm

go 1.22.0

require (
	github.com/9ssi7/txn v0.0.0-00010101000000-000000000000
	github.com/9ssi7/txngorm v0.0.0-00010101000000-000000000000
	github.com/DATA-DOG/go-sqlmock v1.5.2
	gorm.io/driver/postgres v1.5.9
	gorm.io/gorm v1.25.11
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
