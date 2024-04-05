#!/bin/sh

# general
export APP_ENV=local
export APP_PORT=8070

# logging
export LOG_LEVEL=debug

# database
export DB_DIALECT=postgres
export DB_HOST=localhost
export DB_PORT=5432
export DB_NAME=film36exp
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_SILENCE_LOGGER=true