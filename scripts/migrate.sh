#!/bin/bash

export $(cat .env | xargs)

goose -dir ./migrations postgres "$DATABASE_URL_MIGRATIONS" status
goose -dir ./migrations postgres "$DATABASE_URL_MIGRATIONS" up
