module github.com/kheyssper/go-saas-microservice

go 1.22.5

replace github.com/go-saas/sessions => github.com/gorilla/sessions v1.2.1

replace github.com/goava/di => github.com/defval/di v1.12.0

require (
	github.com/jackc/pgx/v4 v4.18.3
	github.com/joho/godotenv v1.5.1
)

require (
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.14.3 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.3 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgtype v1.14.0 // indirect
	github.com/jackc/puddle v1.3.0 // indirect
	golang.org/x/crypto v0.20.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
