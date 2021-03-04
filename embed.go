package main

import "embed"

// content holds our static web server content.
//go:embed pkg/server/ui/build/*
var Content embed.FS
