#!/bin/bash
PACKAGE=$(head -1 go.mod | awk '{print $2}')
echo "USING go.mod PACKAGE : $PACKAGE"
protoc --go_opt=module="${PACKAGE}" --go_out=. --go-grpc_opt=module="${PACKAGE}" --go-grpc_out=. remoteCalculator/calc/calc.proto