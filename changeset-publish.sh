#!/usr/bin/env bash
pnpm exec changeset publish && go run generate.go
