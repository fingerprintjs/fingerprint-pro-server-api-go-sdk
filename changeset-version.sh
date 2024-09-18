#!/usr/bin/env bash
pnpm exec changeset version && bash ./sync-version.sh && go run generate.go
