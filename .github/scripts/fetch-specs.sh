#!/bin/bash
set -e

echo "Fetching latest OpenAPI specs..."

curl -L https://raw.githubusercontent.com/devrev/fern-api-docs/main/fern/apis/public/openapi-public.yaml -o public/spec/openapi-public.yaml

curl -L https://raw.githubusercontent.com/devrev/fern-api-docs/main/fern/apis/beta/openapi-beta.yaml -o beta/spec/openapi-beta.yaml

echo "Specs updated successfully"