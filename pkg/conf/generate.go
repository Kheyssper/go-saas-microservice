package conf

//go:generate protoc --proto_path=../../pkg --proto_path=../../proto --proto_path=../../buf  --go_out=paths=source_relative:../  ../../pkg/conf/conf.proto
