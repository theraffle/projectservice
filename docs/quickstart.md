# Quick Start Guide

This guide leads you to start using `ProjectService`'s gRPC methods .
The contents are as follows.

* [You need](#you-need)
* [Compile Protobuf](#compile-protobuf)
* [Connect to ProjectService](#connect-to-projectservice)


## You need...
- Protoc

## Compile Protobuf
1. Copy `./pb/raffle.proto` into the source repository that you want to use `UserService`.

2. Compile `raffle.proto` according to your source's language.
   ```bash
   # Go Example
    protoc --go_out=plugins=grpc:./src/genproto --go_opt=paths=source_relative ./pb/raffle.proto
   
   # Python Example
    protoc --python_out=plugins=grpc:./src/genproto ./pb/raffle.proto
   ```

## Connect to ProjectService
Default port of projectservice is 7000. Use `projectservice:3550` as address when you create client connection.