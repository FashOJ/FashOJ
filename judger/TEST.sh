#!/bin/bash

set -e
echo "Start Judger..."
echo "Compiler..."
go build -o FashOJ ./cmd/main.go
echo "Complier OK"
echo "Exec..."
echo ""
sudo ./FashOJ
echo ""
echo "Remove exec..."
rm FashOJ
echo "Test OK!"