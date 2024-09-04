#api
# cd desc
# goctl api go -api admin.api -dir ../  --style=goZero --home ../../../../deploy/template

#rpc
# cd proto
# goctl rpc protoc admin.proto --go_out=. --go-grpc_out=. --zrpc_out=../ --style=goZero -m

# sed -i "" 's/,omitempty//g' *.pb.go
