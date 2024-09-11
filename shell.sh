#!/bin/bash

clear

# set environment variables
#export PROJECT_HOME=$(pwd)
#export PATH=$PROJECT_HOME/bin:$PATH

# aliases

# run server in dev mode
alias mad:dev="bash ./scripts/dev.sh"
# build server
alias mad:build="bash ./scripts/build.sh"
# run build
alias mad:run="bash ./scripts/run.sh"

# run client in dev mode
alias mad:client:dev="npm --prefix _client run dev"

# display amount of written go-lines in project
alias mad:stats="bash ./scripts/stats.sh"

# migrations
alias mad:migrate:up="bash ./scripts/migrate-up.sh"
alias mad:migrate:down="bash ./scripts/migrate-down.sh"
alias mad:migrate:create="bash ./scripts/migrate-create.sh"
# migrations shortcuts
alias mad:m:u="mad:migrate:up"
alias mad:m:d="mad:migrate:down"
alias mad:m:c="mad:migrate:create"

#alias ll='ls -la'
#alias gs='git status'
#alias gp='git pull'
#alias gc='git commit -m'

# add any other custom settings or configurations
#export MY_VAR="some_value"

mad:help() {
  echo "\nHelp"
  echo ""
  echo "* Backend"
  printf "%-20s %s\n" "mad:dev" "run server in dev mode"
  printf "%-20s %s\n" "mad:build" "build server"
  printf "%-20s %s\n" "mad:run" "run build"
  echo ""
  echo "* Migrations"
  printf "%-20s %s\n" "mad:migrate:up" "apply new migrations"
  printf "%-20s %s\n" "mad:m:u" ""
  printf "%-20s %s\n" "mad:migrate:down" "rollback (one by one)"
  printf "%-20s %s\n" "mad:m:d" ""
  printf "%-20s %s\n" "mad:migrate:create" "create new .sql migration file"
  printf "%-20s %s\n" "mad:m:c" ""
  echo ""
  echo "* Frontend"
  printf "%-20s %s\n" "mad:client:dev" "run client in dev mode"
  echo ""
  echo "* Other"
  printf "%-20s %s\n" "mad:stats" "display amount of written [go, tsx]-lines in project"
  echo ""
}

echo "MAD Project Environment Loaded"
