# mad

All commands must be executed in project root folder.

## Before start

1. Copy .env.example as .env and update it
2. Use project environment, run: `source ./shell.sh` - it will allow you to use all pre-defined aliases, env variable etc.
2. Run `mad:migrate:up` – this will run migration tool and update database

## Dev Mode

It will build and run server automatically when you change code.

`mad:dev`

## Build

`mad:build`

## Run build

`mad:run`

## Migrations

This project uses goose. Predefined scripts are: 

Up: `mad:migrate:up`

Down: `mad:migrate:down`

Create: `mad:migrate:create`

Other command can be executed by calling directly.