#!/bin/bash

find . -name '*.go' | xargs wc -l | grep -i ' total' | awk '{print $1}'