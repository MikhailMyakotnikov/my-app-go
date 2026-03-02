#!/bin/bash
set -a
. .env
set +a
cd ui-tests/UiTests
dotnet test