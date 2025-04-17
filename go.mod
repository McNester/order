module orders

go 1.24.2

require (
	github.com/go-sql-driver/mysql v1.9.2
	github.com/jmoiron/sqlx v1.4.0
	github.com/joho/godotenv v1.5.1
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	golang.org/x/net v0.39.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)

require (
	cloud_commons v0.0.0
	google.golang.org/grpc v1.71.1
)

require cloud_commons v0.0.0

replace cloud_commons => ./cloud_commons
