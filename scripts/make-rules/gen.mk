# ==============================================================================
# Makefile helper functions for generate necessary files
#

# generate 
ifeq ($(GOOS), darwin)
	wireCmd=xargs -I F sh -c 'cd F && echo && wire'
else
	wireCmd=xargs -i sh -c 'cd {} && echo && wire'
endif

.PHONY: gen.run
# gen.run: gen.wire gen.pb gen.clean 
gen.run: gen.clean gen.pb  

.PHONY: gen.pb
gen.pb: tools.verify.buf tools.verify.gofr
	@echo "===========> Generating pb files *.go from proto file through buf.build"
	@${ROOT_DIR}/scripts/buf.sh
	@gofr wrap grpc server -proto=${ROOT_DIR}/api/grpc/hello/v1/hello.proto
	@gofr wrap grpc client -proto=${ROOT_DIR}/api/grpc/hello/v1/hello.proto

.PHONY: gen.wire
gen.wire: tools.verify.wire
	@echo "===========> Generating wire_gen.go from wire.go file through wire"
	@find internal  -mindepth 2 -maxdepth 2 | grep server | $(wireCmd)

.PHONY: gen.clean
gen.clean:
# @rm -rf ./api/client/{clientset,informers,listers}
	@find ${ROOT_DIR}/api/grpc -type f -name '*.go' -delete
