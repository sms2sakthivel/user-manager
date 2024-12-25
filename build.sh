#!/bin/bash -x

# Step 1: Generate Swagger documentations 
swag init --pdl 3

# Step 2: Build User Service application
go build main.go