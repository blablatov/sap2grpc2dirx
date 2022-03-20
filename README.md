# sap2gRPCdirx
# On Russian

Интеграция систем SAP и Directum RX, в которых нет реализации gRPC. 

Предсталены тестовые модули для обмена данными по принципу REST over gRPC with grpc-gateway for Go.

Схема обмена данными между модулями:

SAP/REST(imitation from web browser) <---> local web-server/sap2grpc <---> gRPC-client/drxclient <---> gRPC-server/drxserver <---> REST/postcreate

После внесения изменений в proto-файл, для генерации нового файла, выполнить в командной строке из директории с файлом drx.proto команду:

protoc drx.proto --go_out=./ --go-grpc_out=./


# On English

Systems integration SAP и Directum RX, they have not gRPC. Test modules REST over gRPC with grpc-gateway for Go.

Sheme exchange of data is like that:

SAP/REST(imitation from web browser) <---> local web-server/sap2grpc <---> gRPC-client/drxclient <---> gRPC-server/drxserver <---> REST/postcreate

After change proto-file, execute command from directory with proto-file: 

protoc drx.proto --go_out=./ --go-grpc_out=./