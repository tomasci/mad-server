# mad

All commands must be executed in project root folder.

## Before start

1. Copy .env.example as .env and update it
2. Run `bash ./scripts/migrate-up.sh` – this will run migration tool

## Dev Mode

It will build and run server automatically when you change code.

`bash ./scripts/dev.sh`

## Build

`bash ./scripts/build.sh`

## Run build

`bash ./scripts/run.sh`

## Migrations

This project uses goose. Predefined scripts are: 

Up: `bash ./scripts/migrate-up.sh`

Down: `bash ./scripts/migrate-down.sh`

Create: `bash ./scripts/migrate-create.sh`

Other command can be executed by calling directly.