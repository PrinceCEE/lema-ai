# Define commands to run api and web
API_DIR := ./api
WEB_DIR := ./web

.PHONY: api web

api:
	cd $(API_DIR) && go run ./cmd/api

web:
	cd $(WEB_DIR) && npm run dev