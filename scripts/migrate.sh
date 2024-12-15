#!/bin/bash

export $(cat .env | xargs)

goose -dir ./migrations postgres "$DATABASE_URL" status
goose -dir ./migrations postgres "$DATABASE_URL" up
