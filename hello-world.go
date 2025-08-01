name: Simple Go Build

on:
	push: 
		branches:
			- main

jobs: 
	build:
		runs-on: ubuntu-latest
		steps:
			- uses: actions/checkout@v3
			- name: setup Go version
			  uses: actions/setup-go@v2
			  with:
			  	go-version: '1.14.0'
			- run: go run hello-world.go