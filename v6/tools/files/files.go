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
)

func main() {
	flag.Parse()
	if *address == "" || *username == "" || *password == "" {
		log.Fatalf("Address, user or password is missing: %s, %s, %s", *address, *username, *password)
	}

	k := kodi.New(*address+"/jsonrpc", *username, *password)

	resp, err := k.Files.GetSources(&kodi.FilesGetSourcesParams{
		Media: kodi.FILES_MEDIA_VIDEO,
	})
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}
	for _, source := range resp.Result.Sources {
		fmt.Println(source.File)
		resp, err := k.Files.GetDirectory(&kodi.FilesGetDirectoryParams{
			Directory: source.File,
			Sort: &kodi.ListSort{
				Method: kodi.SORT_METHOD_TITLE,
			},
		})
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
		for _, f := range resp.Result.Files {
			fmt.Println(f.File)
		}
	}
}
