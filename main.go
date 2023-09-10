package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	_ "github.com/microsoft/go-mssqldb"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.TraceLevel)

	file, err := os.OpenFile(
		"zerologtry.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	defer func() {
		if r := recover(); r != nil {
			log.Panic().Msgf("Panic: %v", r)
		}
	}()
	multi := zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stdout}, file)
	log.Logger = log.Output(multi)

	query := url.Values{}
	query.Add("app name", "MyAppName")

	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword("sa", "password"),
		Host:     fmt.Sprintf("%s:%d", "localhost", 1433),
		RawQuery: query.Encode(),
	}

	db, err := sql.Open("sqlserver", u.String())
	if err != nil {
		panic(err)
	}

	if true {
		log.Error().Msg("ERROR: Its true, so here the error!")
	}

	if 1/3 == 0 {
		log.Warn().Msg("WARN: it's value actually not integer!")
	}

	if complex(0, 1)*complex(0, 1) == -1 {
		log.Info().Msg("INFO: Correctly calculate imaginary number!")
	}

	year, month, day := time.Now().Date()
	log.Debug().Msg(fmt.Sprintf("DEBUG: Now at %d-%d-%d", year, month, day))
	log.Trace().Msg("TRACE: Your random number is " + strconv.Itoa(rand.Int()))

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

// Output:
