package main

import (
	"context"
	"log"
	"web-scrapper-api-fp/configuration"
	"web-scrapper-api-fp/server/rest"
	"web-scrapper-api-fp/services/scrapper"

	"github.com/oklog/run"
	flagd "github.com/open-feature/go-sdk-contrib/providers/flagd/pkg"
	"github.com/open-feature/go-sdk/openfeature"
	"github.com/sethvargo/go-envconfig"
	"github.com/sirupsen/logrus"
)

func main() {

	ctx := context.Background()

	var conf configuration.Configuration
	if err := envconfig.Process(ctx, &conf); err != nil {
		log.Fatal(err)
	}

	openfeature.SetProvider(flagd.NewProvider(flagd.WithHost(conf.OpenFeature.Host), flagd.WithPort(uint16(conf.OpenFeature.Port))))

	// Initialize OpenFeature client
	client := openfeature.NewClient(conf.OpenFeature.Client)

	scrappingService := scrapper.NewScrapper()

	var group run.Group

	group.Add(func() error {
		return rest.Init(conf.RestConfiguration, scrappingService, client)
	}, func(e error) {
		logrus.Fatal(e)
	})

	if err := group.Run(); err != nil {
		logrus.Fatal(err)
	}
}
