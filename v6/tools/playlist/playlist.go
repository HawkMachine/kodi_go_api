package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/HawkMachine/kodi_go_api/v6/kodi"
)

var (
	username = flag.String("u", "", "Kodi username")
	password = flag.String("p", "", "Kodi password")
	address  = flag.String("a", "http://192.168.0.200:9080", "Kodi address")
	id       = flag.Int("id", 0, "Playlist id")

	cmd = flag.String("cmd", "list", "command")
)

func main() {
	flag.Parse()
	if *address == "" || *username == "" || *password == "" {
		log.Fatalf("Address, user or password is missing: %s, %s, %s", *address, *username, *password)
	}

	k := kodi.New(*address+"/jsonrpc", *username, *password)

	switch *cmd {
	case "list":
		response, err := k.Playlist.GetPlaylists()
		if err != nil {
			log.Fatal(err)
		}
		if response.Error != nil {
			log.Fatal(response.Error)
		}
		for idx, playlist := range response.Result {
			fmt.Printf("%d : %3d : %q\n", idx, playlist.PlaylistId, playlist.PlaylistType)
		}
	case "getitems":
		response, err := k.Playlist.GetItems(
			&kodi.PlaylistGetItemsParams{
				PlaylistId: kodi.PlaylistId(*id),
			})
		if err != nil {
			log.Fatal(err)
		}
		if response.Error != nil {
			log.Fatal(response.Error)
		}
		for idx, item := range response.Result.Items {
			fmt.Printf("%d : %#v\n", idx, item)
		}
	default:
		log.Fatal("Unrecognized command: " + *cmd)
	}
}
