#!/bin/bash
set -e

# Argument: path to .env-file, .env.dev by default
ENV_FILE="${1:-.env.dev}"

if [ ! -f "$ENV_FILE" ]; then
    echo "Env file '$ENV_FILE' not found!"
    exit 1
fi

echo "Loading environment from $ENV_FILE..."
set -a
. "$ENV_FILE"
set +a

echo "Running UI tests..."
cd ui-tests/UiTests
dotnet test