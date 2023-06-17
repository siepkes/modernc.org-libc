# Copyright 2023 The Libc Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	all clean edit editor generate generate2 test

all:

clean:
	rm -f log-* cpu.test mem.test *.out
	go clean

edit:
	@touch log
	@if [ -f "Session.vim" ]; then gvim -S & else gvim -p Makefile *.go & fi

editor:
	gofmt -l -s -w *.go 2>&1 | tee log-editor
	go test 2>&1 | tee -a log-editor
	go install -v  2>&1 | tee -a log-editor
	go build -o /dev/null generator*.go
	staticcheck 2>&1 | tee -a log-editor
	golint 2>&1 | tee -a log-editor

generate:
	go run generator*.go 2> log-generate-errors | tee log-generate
	cat log-generate-errors

generate2:
	go run generator*.go 2>&1 | tee log-generate

test:
	go test -v -timeout 24h 2>&1 | tee log-test
	grep -a 'TOTAL\|PASS' log-test || true 2>&1 | tee -a log-test
