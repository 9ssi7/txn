module github.com/9ssi7/txn/txnredis

go 1.22.0

replace github.com/9ssi7/txn/tx => ../tx

require (
	github.com/9ssi7/txn/tx v0.0.0-00010101000000-000000000000
	github.com/go-redis/redismock/v9 v9.2.0
	github.com/redis/go-redis/v9 v9.6.0
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
