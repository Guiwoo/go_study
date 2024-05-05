module nats

go 1.20

replace github.com/guiwoopark/socket => ../socket

require github.com/nats-io/nats.go v1.34.1

require (
	github.com/klauspost/compress v1.17.7 // indirect
	github.com/nats-io/nkeys v0.4.7 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
)
