# fabux
================
This package is being developed

Only a little bit of functionality from now

[![Build Status](https://travis-ci.org/benlightning/fabux.svg)](https://travis-ci.org/benlightning/fabux/fabcore) [![Go Walker](http://gowalker.org/api/v1/badge)](http://gowalker.org/github.com/benlightning/fabux)

This is an open source project about ssh connection tool for the Go programming language.

Development based on **go 1.5.1**

## Installation

Simple as it takes to type the following command:

    go get github.com/benlightning/fabux

Dependency:
	
	[pscp](https://the.earth.li/~sgtatham/putty/latest/x86/pscp.exe) windows
	scp
	github.com/Unknwon/goconfig

## Usage

~~~
Usage: fabux [options...]

Options:
  -g 
~~~

## Contribute

Your contribute is welcome, but you have to check following steps after you added some functions and commit them:

1. Make sure you wrote user-friendly comments for **all functions** .
2. Make sure you wrote test cases with any possible condition for **all functions** in file `*_test.go`.
3. Make sure you wrote benchmarks for **all functions** in file `*_test.go`.
4. Make sure you ran `go test` and got **PASS** .

## License

Copyright 2014 Google Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.