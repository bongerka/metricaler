package main

import (
	"github.com/bongerka/metricaler/internal/app"
	"github.com/bongerka/metricaler/internal/config"
)

func main() {
	var cfg config.Config
	config.MustParse(&cfg)

	app.NewApp(&cfg).Run()
}
