#!/bin/bash

mockgen \
  -destination pkg/mock_example/mock_thing.go \
  -package mock_example \
  github.com/briantkennedy/mockbug/pkg/example Thing
