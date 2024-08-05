gen-protoc:
	cd ./order && protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/service.proto
	cd ..
	cd ./payment && protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/service.proto
prepare-env:
	docker compose up -d

shutdown-env:
	docker compose down