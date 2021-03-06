#!/bin/bash
# wait-for-postgres.sh

set -e

host="$1"
shift
cmd="$@"
echo $PGPASSWORD
until PGPASSWORD=$POSTGRES_PASSWORD psql -h "$host" -U "postgres" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done
