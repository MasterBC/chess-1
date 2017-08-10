#!/bin/sh
govendor update chess/common/consul
govendor update chess/common/define
govendor update chess/common/log
govendor update chess/common/services
go run *.go --tcp-listen :8898 --ws-listen :8899