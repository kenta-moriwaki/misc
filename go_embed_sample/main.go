package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed data/chunked/*
var dataFS embed.FS

type Article struct {
	ID    uint32 `json:"id"`
	Guid  string `json:"guid"`
	Title string `json:"title"`
	Text  string `json:"text"`
	// CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt time.Time `json:"updatedAt"`
	Tags []string `json:"tags"`
}

func main() {
	articles := map[uint32]Article{}
	tagIndex := map[string]map[uint32]struct{}{}

	chunkedDir, err := fs.Sub(dataFS, "data/chunked")
	if err != nil {
		panic(err)
	}

	paths, err := fs.Glob(chunkedDir, "*.json")
	if err != nil {
		panic(err)
	}

	for _, p := range paths {
		func() {
			f, err := chunkedDir.Open(p)
			if err != nil {
				panic(err)
			}
			defer f.Close()

			b, _ := ioutil.ReadAll(f)

			var article Article
			json.Unmarshal(b, &article)

			articles[article.ID] = article
			for _, t := range article.Tags {
				if m, ok := tagIndex[t]; ok {
					m[article.ID] = struct{}{}
				} else {
					tagIndex[t] = map[uint32]struct{}{}
					tagIndex[t][article.ID] = struct{}{}
				}
			}
		}()
	}

	fmt.Printf("%+v\n", tagIndex["consequat"])
	fmt.Printf("%+v\n", tagIndex["officia"])
	fmt.Printf("%+v\n", tagIndex["nostrud"])
	fmt.Printf("%+v\n", tagIndex["hoge"])
}
