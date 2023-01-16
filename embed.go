package main

import "embed"

//go:embed all:frontend/build
var FrontendFs embed.FS
