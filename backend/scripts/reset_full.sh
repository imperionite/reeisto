#!/bin/bash
set -e

echo "Full reset..."

./scripts/reset_data.sh
./scripts/run_migrations.sh

echo "Full reset complete."