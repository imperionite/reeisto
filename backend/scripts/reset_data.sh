#!/bin/bash
set -e

source .env

echo "Resetting data..."

PGPASSWORD=$DB_PASSWORD psql \
  -h "$DB_HOST" \
  -p "$DB_PORT" \
  -U "$DB_USER" \
  -d "$DB_NAME" \
  -c "TRUNCATE TABLE inventories, rees, transactions, users RESTART IDENTITY CASCADE;"

echo "Data reset complete."