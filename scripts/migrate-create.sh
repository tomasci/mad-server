#!/bin/bash

. ./scripts/env.sh

# Prompt the user for the migration name
read -p "Enter the migration name: " migration_name

goose -dir ./migrations create "$migration_name" sql