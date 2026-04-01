#!/bin/bash
set -e

# ----------------------------
# Load .env safely
# ----------------------------
if [ -f .env ]; then
  set -o allexport
  source .env
  set +o allexport
fi

# ----------------------------
# Validate required variables
# ----------------------------
: "${DB_HOST:?Need to set DB_HOST}"
: "${DB_PORT:?Need to set DB_PORT}"
: "${DB_NAME:?Need to set DB_NAME}"
: "${DB_USER:?Need to set DB_USER}"
: "${DB_PASSWORD:?Need to set DB_PASSWORD}"

echo "Checking database connection..."

PGPASSWORD=$DB_PASSWORD psql \
  -h "$DB_HOST" \
  -p "$DB_PORT" \
  -U "$DB_USER" \
  -d "$DB_NAME" \
  -c "\q"

echo "Connection OK."

# ----------------------------
# Run migrations
# ----------------------------
for f in migrations/*.sql; do
  echo "Running migration: $f"
  PGPASSWORD=$DB_PASSWORD psql \
    -h "$DB_HOST" \
    -p "$DB_PORT" \
    -U "$DB_USER" \
    -d "$DB_NAME" \
    -f "$f"
done

echo "All migrations executed."