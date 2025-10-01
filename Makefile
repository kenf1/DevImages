.PHONY: install_dep run_zap clean compile grm buildct lint

install_dep: #Install zap
	cd ZigDevEx && \
	zig fetch --save "git+https://github.com/zigzap/zap#v0.9.1"

run_zap: #Build & run zap
	cd ZigDevEx && \
	zig build run

compile: #Compile create_template
	cd create_template && \
	deno compile main.ts && \
	mv create_template ../CreateTemplate

grm: #Run main.go
	cd templater && go run *.go

buildct: #Build binary
	cd templater && go build

lint: #Lint Go
	cd templater && golangci-lint run

fmt: #Format python
	ruff format
	find . -type d -name ".ruff_cache" | xargs rm -rf

clean: fmt #Cleanup entire repo
	find . -type d -name "zig-out" | xargs rm -rf
	find . -type d -name ".zig-cache" | xargs rm -rf
	find . -type d -name "__pycache__" | xargs rm -rf
	find . -type d -name ".pytest_cache" | xargs rm -rf