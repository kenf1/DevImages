help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} \
	/^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } \
	/^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) }' $(MAKEFILE_LIST)

install_dep: ##Install zap
	zig fetch --save "git+https://github.com/zigzap/zap#v0.9.1"

runzap: ##Build & run zap
	zig build run

cleanzig: ##Rm all cache
	find . -type d -name "zig-out" | xargs rm -rf
	find . -type d -name ".zig-cache" | xargs rm -rf