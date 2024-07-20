module github.com/9ssi7/txn

go 1.22.0

replace (
	github.com/9ssi7/txn/tx v0.0.1-beta.5 => ./tx
	github.com/9ssi7/txn/txngorm v0.0.1-beta.5 => ./txngorm
	github.com/9ssi7/txn/txnmongo v0.0.1-beta.5 => ./txnmongo
	github.com/9ssi7/txn/txnredis v0.0.1-beta.5 => ./txnredis
)
