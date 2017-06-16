GO := $(shell which go)
GO_VERSION := $(shell $(GO) version)
GOAPP := $(shell which goapp)

CARD_REFUND_ROOT := $(shell echo "${GOPATH}/src/github.com/MarcGrol/cardRefund")

#
# Configation for production-environment
#
PROD_APP_NAME := "duxxieplatformprod"
PROD_APP_VERSION := $(shell date +"%Y-%m-%d")
PROD_APP_HJID := "361317"
PROD_APP_SENTRY_DSN := "https://14ca47a5f6974877ad4c74df3577ec4a@sentry.io/122796"
PROD_SECRET_KEY_FILE := "duxxieplatformprod-clientsecret.json"
PROD_ANALYTICS_UID := ""
PROD_GTM_CONTAINER_ID := "GTM-WGCXPFP"

all: gen test install

help:
	@echo "\tdeps: installs all dependencies"
	@echo "\tgen: generates boilerplate code"
	@echo "\tfrontend: Run all frontend tests"
	@echo "\tbackend: Run all backend tests"
	@echo "\ttest: Run all tests"
	@echo "\trun: Run application locally in dev mode"
	@echo "\ttestdeploy: Deploy on test environment"
	@echo "\tauditdeploy: Deploy on audit environment"
	@echo "\tdemodeploy, demodeploy: Deploy on demo environment"
	@echo "\tproddeploy, proddeploy: Deploy on production environment"

checkmaster:
	@echo "---------------------------"
	@echo "  FROM MASTER BRANCH ONLY  "
	@echo "Performing dependency check"
	@echo "---------------------------"
	$(GO) get -u -t github.com/MarcGrol/cardRefund/...

deps:
	@echo "---------------------------"
	@echo "Performing dependency check"
	@echo "---------------------------"
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/kardianos/govendor
	go get -u google.golang.org/appengine/...
	go get -u github.com/gorilla/context
	govendor sync -v                                    # populate vendor directory with specified version of golangAnnotations
	govendor install +vendor                            # make sure we get the vendored version of golangAnnotations
	go get -t ./...                                     # get the application with all its deps
	$(CARD_REFUND_ROOT)/scripts/get_appengine.sh    # Need appengine to run goapp tests and to deploy
	$(CARD_REFUND_ROOT)/scripts/get_gcloud.sh       # Need gcloud to perform deployment
	(cd emails; yarn install; yarn run build;)          # build email templates (used in backend)
	(cd frontend; yarn install;)                        # prepare environment for frontend application

verify:
	@echo "----------------------------"
	@echo "Run static analysis on source-code"
	@echo "----------------------------"
	(cd backend; $(GO) vet ./...;)
	#golint ./...

generate:
	@echo "----------------------"
	@echo "Generating source-code"
	@echo "----------------------"
	$(GO) generate ./...

imports:
	@echo "------------------"
	@echo "Optimizing imports"
	@echo "------------------"
	for i in `find . -name "*.go"`; do goimports -w -local github.com/ $${i}; done

format:
	@echo "----------------------"
	@echo "Formatting source-code"
	@echo "----------------------"
	for i in `find . -name "*.go"`; do gofmt -s -w $${i}; done

gen: generate imports format

test: backend

backend: gen backendtest

backendtest:
	@echo "-------------"
	@echo "Running backend tests"
	@echo "-------------"
	$(GOAPP) test ./...                     # run appengine tests
	$(GO) test ./...                        # run unit tests
	make format

coverage:
	@echo "-------------"
	@echo "Running coverage"
	@echo "-------------"
	$(CARD_REFUND_ROOT)/scripts/coverage.sh --html

install:
	@echo "----------------"
	@echo "Installing for $(GO_VERSION)"
	@echo "----------------"
	govendor install +vendor ./...
	$(GO) install ./...

clean:
	find . -name "\$$ast.json" -exec rm {} \;
	find . -name "\$$httpService.go" -exec rm {} \;
	find . -name "\$$httpTestService.go" -exec rm {} \;
	find . -name "\$$httpServiceHelpers_test.go" -exec rm {} \;
	find . -name "\$$httpClientForService.go" -exec rm {} \;
	find . -name "\$$eventHandler.go" -exec rm {} \;
	make gen
	$(GO) clean ./...

run:
	@echo "------------------------------------------"
	@echo "Running $(APP_NAME):$(APP_VERSION) locally"
	@echo "------------------------------------------"
	$(GOAPP) serve -host=0.0.0.0 $(CARD_REFUND_ROOT)/main/

e2erun:
	@echo "------------------------------------------"
	@echo "Running $(APP_NAME):$(APP_VERSION) locally"
	@echo "------------------------------------------"
	dev_appserver.py --clear_datastore --port=8888 --env_var DEV_E2E_MODE=true --skip_sdk_update_check true $(CARD_REFUND_ROOT)/main/app.yaml

clearrun:
	@echo "------------------------------------------"
	@echo "Running $(APP_NAME):$(APP_VERSION) locally"
	@echo "------------------------------------------"
	$(GOAPP) serve -clear_datastore -host=0.0.0.0 $(CARD_REFUND_ROOT)/main/

deploy:
	@echo "---------------------------------------------------------------"
	@echo "Deploying interactively $(APP_NAME):$(APP_VERSION) to the cloud"
	@echo "---------------------------------------------------------------"
	./scripts/interactivedeploy.sh $(APP_NAME) $(APP_VERSION)

travisdeploy:
	@echo "---------------------------------------------------------------"
	@echo "Deploying automated $(APP_NAME):$(APP_VERSION) to the cloud"
	@echo "---------------------------------------------------------------"
	./scripts/travisdeploy.sh $(APP_NAME) $(APP_VERSION)

prodvars:
	$(eval APP_NAME := $(PROD_APP_NAME))
	$(eval APP_VERSION := $(PROD_APP_VERSION))
	$(eval HJID := $(PROD_APP_HJID))
	$(eval SENTRY_DSN := $(PROD_APP_SENTRY_DSN))
	$(eval SECRET_KEY_FILE := $(PROD_SECRET_KEY_FILE))
	$(eval ANALYTICS_UID := $(PROD_ANALYTICS_UID))
	$(eval GTM_CONTAINER_ID := $(PROD_GTM_CONTAINER_ID))

prodbuild: prodvars frontend backend

proddeploy: prodbuild travisdeploy

.PHONY:
	help deps verify gen test backendtest backend coverage install clean run e2erun \
	clearrun all \
	travisdeploy deploy \
	prodvars prodbuild proddeploy
