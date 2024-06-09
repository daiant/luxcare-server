module luxcare/server

go 1.22.4

replace luxcare/contact => ./contact

require (
	luxcare/contact v0.0.0-00010101000000-000000000000
	luxcare/database v0.0.0-00010101000000-000000000000
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
)

replace luxcare/database => ./database
