package main

import (
	//"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/HawkMachine/kodi_go_api/v6/kodi"
)

var (
	username = flag.String("u", "", "Kodi username")
	password = flag.String("p", "", "Kodi password")
	duration = flag.Duration("d", 7*24*time.Hour, "Duration")
	address  = flag.String("a", "http://192.168.0.200:9080", "Kodi address")
)

func main() {
	flag.Parse()
	if *address == "" || *username == "" || *password == "" {
		log.Fatalf("Address, user or password is missing: %s, %s, %s", *address, *username, *password)
	}

	k := kodi.New(*address+"/jsonrpc", *username, *password)

	sinceDate := time.Now().Add(-*duration)

	epResult, err := k.VideoLibrary.GetEpisodes(&kodi.VideoLibraryGetEpisodesParams{
		Properties: []kodi.VideoFieldsEpisode{
			kodi.EPISODE_FIELD_SHOW_TITLE,
			kodi.EPISODE_FIELD_TITLE,
			kodi.EPISODE_FIELD_LAST_PLAYED,
		},
		Filter: &kodi.VideoLibraryGetEpisodesFilter{
			ListFilterEpisodes: &kodi.ListFilterEpisodes{
				ListFilterRuleEpisodes: &kodi.ListFilterRuleEpisodes{
					Field:    kodi.EPISODE_FILTER_FIELD_LAST_PLAYED,
					Operator: kodi.OPERATOR_AFTER,
					Value:    sinceDate.Format("2006-01-02"),
				},
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	byDate := map[time.Time]int{}
	for _, edetails := range epResult.Result.Episodes {
		t, err := time.Parse("2006-01-02 15:04:05", edetails.LastPlayed)
		if err != nil {
			log.Fatal(err)
		}
		year, month, day := t.Date()
		d := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		v, ok := byDate[d]
		if ok {
			byDate[d] = v + 1
		} else {
			byDate[d] = v + 1
		}
	}
	for s, v := range byDate {
		fmt.Printf("%30s : %d\n", s, v)
	}
}
