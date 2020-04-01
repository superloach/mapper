package main

import (
	"github.com/superloach/minori"

	"github.com/superloach/snart/bot"
	"github.com/superloach/snart/errs"
	"github.com/superloach/snart/plugin"
	"github.com/superloach/snart/route"
)

var Log *minori.Logger

func Register(name string, b *bot.Bot) error {
	_f := "Register"

	Log = plugin.Log.GetLogger(name)
	Log.Info(_f, "forking registration")

	go func() {
		b.DB.Easy(POIDB)
		b.DB.Easy(POITable)

		err := routes(name, b)
		if err != nil {
			errs.Wrap(&err, `routes(%#v, b)`, name)
			Log.Warn(_f, err)
			return
		}
	}()

	Log.Info(_f, "forked registration")
	return nil
}

func routes(name string, b *bot.Bot) error {
	_f := "routes"
	Log.Info(_f, "registering routes")

	poi := Poi(b.DB)

	b.AddRoute(
		&route.Route{
			Name:  "pois",
			Match: "pois?",
			Desc:  "Search for any POIs. (Alias: `poi`)",
			Cat:   name,
			Okay:  nil,
			Func:  poi,
		},
		&route.Route{
			Name:  "gyms",
			Match: "g(yms?)?",
			Desc:  "Search for Pokemon Go gyms. (Alias: `gym`, `g`)",
			Cat:   name,
			Okay:  nil,
			Func:  poi,
		},
		&route.Route{
			Name:  "stops",
			Match: "s(tops?)?",
			Desc:  "Search for Pokemon Go stops. (Alias: `stop`, `s`)",
			Cat:   name,
			Okay:  nil,
			Func:  poi,
		},
	)

	Log.Info(_f, "registered routes")
	return nil
}