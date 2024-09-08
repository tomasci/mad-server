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

echo "MAD Project Environment Loaded"
