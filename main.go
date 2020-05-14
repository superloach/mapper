package main

import (
	"fmt"

	"github.com/namsral/flag"
	"github.com/superloach/minori"

	"github.com/go-snart/bot"
	"github.com/go-snart/db"

	_ "github.com/go-snart/plugin-admin"
	_ "github.com/go-snart/plugin-help"

	_ "github.com/superloach/mapper/plugin"
)

var (
	debug = flag.Bool("debug", false, "print debug messages")

	dbhost = flag.String("dbhost", "localhost", "rethinkdb host")
	dbport = flag.Int("dbport", 28015, "rethinkdb port")
	dbuser = flag.String("dbuser", "admin", "rethinkdb username")
	dbpass = flag.String("dbpass", "", "rethinkdb password")
)

var Log = minori.GetLogger("mapper")

func main() {
	_f := "main"
	flag.Parse()

	if *debug {
		minori.Level = minori.DEBUG
		Log.Debug(_f, "debug mode")
	} else {
		minori.Level = minori.INFO
	}

	d := &db.DB{
		Host: *dbhost,
		Port: *dbport,
		User: *dbuser,
		Pass: *dbpass,
	}

	// make bot
	b, err := bot.MkBot(d)
	if err != nil {
		err = fmt.Errorf("mkbot %#v: %w", d, err)
		Log.Fatal(_f, err)
	}

	// run the bot
	err = b.Start()
	if err != nil {
		err = fmt.Errorf("start: %w", err)
		Log.Fatal(_f, err)
	}

	Log.Info(_f, "bye!")
}
