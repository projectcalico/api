PACKAGE_NAME    ?= github.com/projectcalico/api
GO_BUILD_VER    ?= v0.53
GOMOD_VENDOR    := false
GIT_USE_SSH      = true
LOCAL_CHECKS     = lint-cache-dir goimports check-copyright

BINDIR ?= bin
BUILD_DIR ?= build
TOP_SRC_DIRS = pkg

##############################################################################
# Download and include Makefile.common before anything else
#   Additions to EXTRA_DOCKER_ARGS need to happen before the include since
#   that variable is evaluated when we declare DOCKER_RUN and siblings.
##############################################################################
MAKE_BRANCH?=$(GO_BUILD_VER)
MAKE_REPO?=https://raw.githubusercontent.com/projectcalico/go-build/$(MAKE_BRANCH)

Makefile.common: Makefile.common.$(MAKE_BRANCH)
	cp "$<" "$@"
Makefile.common.$(MAKE_BRANCH):
	# Clean up any files downloaded from other branches so they don't accumulate.
	rm -f Makefile.common.*
	curl --fail $(MAKE_REPO)/Makefile.common -o "$@"

include Makefile.common

build: gen-files examples

###############################################################################
# This section contains the code generation stuff
###############################################################################
.generate_execs: lint-cache-dir\
	$(BINDIR)/defaulter-gen \
	$(BINDIR)/deepcopy-gen \
	$(BINDIR)/conversion-gen \
	$(BINDIR)/client-gen \
	$(BINDIR)/lister-gen \
	$(BINDIR)/informer-gen \
	$(BINDIR)/openapi-gen
	touch $@

$(BINDIR)/deepcopy-gen:
	$(DOCKER_GO_BUILD) sh -c "GOBIN=/go/src/$(PACKAGE_NAME)/$(BINDIR) go install k8s.io/code-generator/cmd/deepcopy-gen"

$(BINDIR)/client-gen:
	$(DOCKER_GO_BUILD) sh -c "GOBIN=/go/src/$(PACKAGE_NAME)/$(BINDIR) go install k8s.io/code-generator/cmd/client-gen"

$(BINDIR)/lister-gen:
	$(DOCKER_GO_BUILD) sh -c "GOBIN=/go/src/$(PACKAGE_NAME)/$(BINDIR) go install k8s.io/code-generator/cmd/lister-gen"

$(BINDIR)/informer-gen:
	$(DOCKER_GO_BUILD) sh -c "GOBIN=/go/src/$(PACKAGE_NAME)/$(BINDIR) go install k8s.io/code-generator/cmd/informer-gen"

$(BINDIR)/defaulter-gen:
	$(DOCKER_GO_BUILD) sh -c "GOBIN=/go/src/$(PACKAGE_NAME)/$(BINDIR) go install k8s.io/code-generator/cmd/defaulter-gen"

$(BINDIR)/conversion-gen:
	$(DOCKER_GO_BUILD) sh -c "GOBIN=/go/src/$(PACKAGE_NAME)/$(BINDIR) go install k8s.io/code-generator/cmd/conversion-gen"

$(BINDIR)/openapi-gen:
	$(DOCKER_GO_BUILD) sh -c "GOBIN=/go/src/$(PACKAGE_NAME)/$(BINDIR) go install k8s.io/code-generator/cmd/openapi-gen"

# Regenerate all files if the gen exes changed or any "types.go" files changed
.PHONY: gen-files
gen-files .generate_files: lint-cache-dir .generate_execs clean-generated
	# Generate defaults
	$(DOCKER_RUN) $(CALICO_BUILD) \
	   sh -c '$(GIT_CONFIG_SSH) $(BINDIR)/defaulter-gen \
		--v 1 --logtostderr \
		--go-header-file "/go/src/$(PACKAGE_NAME)/hack/boilerplate/boilerplate.go.txt" \
		--input-dirs "$(PACKAGE_NAME)/pkg/apis/projectcalico/v3" \
		--extra-peer-dirs "$(PACKAGE_NAME)/pkg/apis/projectcalico/v3" \
		--output-file-base "zz_generated.defaults"'
	# Generate deep copies
	$(DOCKER_RUN) $(CALICO_BUILD) \
	   sh -c '$(GIT_CONFIG_SSH) $(BINDIR)/deepcopy-gen \
		--v 1 --logtostderr \
		--go-header-file "/go/src/$(PACKAGE_NAME)/hack/boilerplate/boilerplate.go.txt" \
		--input-dirs "$(PACKAGE_NAME)/pkg/apis/projectcalico/v3" \
		--bounding-dirs $(PACKAGE_NAME) \
		--output-file-base zz_generated.deepcopy'

	# generate all pkg/client contents
	$(DOCKER_RUN) $(CALICO_BUILD) \
	   sh -c '$(GIT_CONFIG_SSH) $(BUILD_DIR)/update-client-gen.sh'

	# generate openapi
	$(DOCKER_RUN) $(CALICO_BUILD) \
	   sh -c '$(GIT_CONFIG_SSH) $(BINDIR)/openapi-gen \
		--v 1 --logtostderr \
		--go-header-file "/go/src/$(PACKAGE_NAME)/hack/boilerplate/boilerplate.go.txt" \
		--input-dirs "$(PACKAGE_NAME)/pkg/apis/projectcalico/v3,k8s.io/api/core/v1,k8s.io/api/networking/v1,k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/apimachinery/pkg/version,k8s.io/apimachinery/pkg/runtime,k8s.io/apimachinery/pkg/util/intstr,$(PACKAGE_NAME)/pkg/lib/numorstring" \
		--output-package "$(PACKAGE_NAME)/pkg/openapi"'
	$(DOCKER_GO_BUILD) \
           sh -c '$(BINDIR)/openapi-gen \
                --v 1 --logtostderr \
                --go-header-file "/go/src/$(PACKAGE_NAME)/hack/boilerplate/boilerplate.go.txt" \
                --input-dirs "$(PACKAGE_NAME)/pkg/lib/numorstring" \
                --output-package "$(PACKAGE_NAME)/pkg/lib/numorstring"; \
                sed -i "/numorstring /d" ./pkg/lib/numorstring/openapi_generated.go'
                # Above 'sed' to workaround a bug in openapi-gen which ends up
                # importing "numorstring github.com/.../lib/numorstring" causing eventual build error
	touch .generate_files
	$(MAKE) fix


.PHONY: lint-cache-dir
lint-cache-dir:
	mkdir -p $(CURDIR)/.lint-cache

.PHONY: check-copyright
check-copyright:
	@hack/check-copyright.sh

.PHONY: clean
clean: clean-bin
	rm -rf .lint-cache Makefile.common*

clean-generated:
	rm -f .generate_files
	find $(TOP_SRC_DIRS) -name zz_generated* -exec rm {} \;
	# rollback changes to the generated clientset directories
	# find $(TOP_SRC_DIRS) -type d -name *_generated -exec rm -rf {} \;
	rm -rf pkg/client/clientset_generated pkg/client/informers_generated pkg/client/listers_generated pkg/openapi

clean-bin:
	rm -rf $(BINDIR) \
	    .generate_execs \

.PHONY: examples
examples: bin/list-gnp

bin/list-gnp: examples/list-gnp/main.go
	@echo Building list-gnp example binary...
	mkdir -p bin
	$(DOCKER_GO_BUILD) sh -c '$(GIT_CONFIG_SSH) \
	   	go build -v -o $@ -v $(LDFLAGS) "examples/list-gnp/main.go"' 

WHAT?=.
GINKGO_FOCUS?=.*

.PHONY:ut
ut:
	$(DOCKER_RUN) --privileged $(CALICO_BUILD) \
		sh -c 'cd /go/src/$(PACKAGE_NAME) && ginkgo -r -focus="$(GINKGO_FOCUS)" $(WHAT)'