#!/usr/bin/env bash
go run generate.go && pnpm exec changeset publish
