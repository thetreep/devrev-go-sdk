#!/bin/bash
set -e

echo "Fetching latest changelogs..."

mkdir -p changelogs/public changelogs/beta

git clone --depth 1 --sparse https://github.com/devrev/fern-api-docs.git temp_repo
cd temp_repo
git sparse-checkout init --cone
git sparse-checkout set fern/apis/public/changelog fern/apis/beta/changelog

cp -r fern/apis/public/changelog/* ../changelogs/public/
cp -r fern/apis/beta/changelog/* ../changelogs/beta/

cd ..
rm -rf temp_repo

echo "Changelogs updated successfully"