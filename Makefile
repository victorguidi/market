GO := go
GOBUILD := $(GO) build
GOFLAGS := -v
BACKEND_DIR := ./src
# FRONTEND_DIR := ./frontend

# Define the build targets
# all:check create-db create-cert frontend-start backend-build
# all:check create-db create-cert backend-build
all:check create-db backend-build

# Test if golang and nodejs are installed
check:
	@which go || (echo "Go is not installed" && exit 1)

#Create the db file
create-db:
	@test -d $(BACKEND_DIR)/databases/ || mkdir $(BACKEND_DIR)/databases/
	@test -f $(BACKEND_DIR)/databases/market.db || touch $(BACKEND_DIR)/databases/market.db

#Create the db file
# create-cert:
# 	@test -d $(BACKEND_DIR)/selfCertificate || ./certgen.sh

# frontend-start:
# 	@cd $(FRONTEND_DIR) && pnpm install
# 	@cd $(FRONTEND_DIR) && pnpm run build
# 	# @cd $(FRONTEND_DIR) && npm run build
# 	@test -d $(BACKEND_DIR)/static || mkdir $(BACKEND_DIR)/static
# 	@cp -R $(FRONTEND_DIR)/static/* $(BACKEND_DIR)/static

backend-build:
	@cd $(BACKEND_DIR) && $(GOBUILD) $(GOFLAGS) -o market .

run: all
	@cd $(BACKEND_DIR) && ./market

clean:
	@rm -rf $(BACKEND_DIR)/bin/*

.PHONY: all create-db run clean check backend-build
