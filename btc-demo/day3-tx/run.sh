#!/bin/bash
rm block
rm blockChain.db
rm blockChain.db.lock
go build -o block *.go
./block