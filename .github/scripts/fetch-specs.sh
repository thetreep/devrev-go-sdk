#!/bin/bash
set -e

echo "Fetching latest OpenAPI specs..."

curl -fL https://developer.devrev.ai/public/openapi/public.yaml -o public/spec/openapi-public.yaml

curl -fL https://developer.devrev.ai/openapi/beta.yaml -o beta/spec/openapi-beta.yaml

echo "Specs updated successfully"