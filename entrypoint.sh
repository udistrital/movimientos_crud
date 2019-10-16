#!/usr/bin/env bash

set -e
set -u
set -o pipefail

if [ -n "${PARAMETER_STORE:-}" ]; then
  export MOVIMIENTOS_CRUD_DB_USER="$(aws ssm get-parameter --name /${PARAMETER_STORE}/movimientos_crud/db/username --output text --query Parameter.Value)"
  export MOVIMIENTOS_CRUD_DB_PASSWORD="$(aws ssm get-parameter --with-decryption --name /${PARAMETER_STORE}/movimientos_crud/db/password --output text --query Parameter.Value)"
fi

exec ./main "$@"
