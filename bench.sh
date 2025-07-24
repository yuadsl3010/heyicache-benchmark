#!/bin/bash

go test -bench=. -benchmem -benchtime=10s
