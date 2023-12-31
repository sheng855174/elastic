PWD=$(shell pwd)
BEAT_TYPE?=beat
BEAT_NAME?=beatpath/test${BEAT_TYPE}
BEAT_PATH=${GOPATH}/src/${BEAT_NAME}
ES_BEATS=${GOPATH}/src/github.com/elastic/beats
PREPARE_COMMAND?=

-include ${ES_BEATS}/dev-tools/make/mage-install.mk

# Runs test build for mock beat
.PHONY: test
test: prepare-test
	cd ${BEAT_PATH} ; \
	export PATH=$${GOPATH}/bin:$${PATH}; \
	git config user.email "beats-jenkins@test.com" || exit 1 ; \
	git config user.name "beats-jenkins" || exit 1 ; \
	$(MAKE) check CHECK_HEADERS_DISABLED=y || exit 1 ; \
	$(MAKE) || exit 1 ; \
	mage test

.PHONY: test-package
test-package: test
	cd ${BEAT_PATH} ; \
	export PATH=$${GOPATH}/bin:$${PATH}; \
	mage package

.PHONY: prepare-test
prepare-test:: mage
	rm -fr ${BEAT_PATH}
	mkdir -p ${BEAT_PATH}
	cd ${GOPATH}/src/github.com/elastic/beats/ ; \
	export MODULE=elastic ; \
	export METRICSET=test ; \
	export NEWBEAT_PROJECT_NAME=Testbeat ; \
	export NEWBEAT_GITHUB_NAME=ruflin ; \
	export NEWBEAT_BEAT_PATH=${BEAT_NAME} ; \
	export NEWBEAT_FULL_NAME="Nicolas Ruflin" ; \
	export NEWBEAT_TYPE=${BEAT_TYPE} ; \
	export NEWBEAT_DEV=1 ; \
	export NEWBEAT_BEATS_REVISION=$(shell git rev-parse origin/7.16) ; \
	mage GenerateCustomBeat

# Runs test build for the created beat
.PHONY: test-build
test-build: test
	cd ${BEAT_PATH}/dev-tools/packer ; \
	$(MAKE) deps ; \
	$(MAKE) images

# Cleans up environment
.PHONY: clean
clean:
	rm -rf ${BEAT_PATH}
