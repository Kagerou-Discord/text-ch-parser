package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hclparse"
)

type Channel struct {
	Type                  string  `hcl:",label"`
	ResourceName          string  `hcl:",label"`
	Name                  string  `hcl:"name"`
	ServerId              *int    `hcl:"server_id"`
	Position              *int    `hcl:"position"`
	Category              *string `hcl:"category"`
	SyncPermsWithCategory *bool   `hcl:"sync_perms_with_category"`

	// discord_text_channel
	Topic *string `hcl:"topic"`
	Nsfw  *bool   `hcl:"nsfw"`

	// discord_voice_channel
	Bitrate   *int `hcl:"bitrate"`
	UserLimit *int `hcl:"user_limit"`
}

type Config struct {
	Resource []Channel `hcl:"resource,block"`
}

const usages = `
Usage:
text-ch-parser [flags]
The flags are:
-f: the HCL file name you want to parse (default: main.tf)
-w: the Markdown file name you want to write parsed HCL (default: main.md)
-d: whether outputs logs for debugging (default: false)`

func flagUsage() {
	log.Fatalln(usages)
}

func main() {
	flag.Usage = flagUsage
	hclFile, writeFile, isDebug := parseArgs()

	config, parseErr := parseHcl(hclFile)
	if parseErr != "" {
		log.Fatal(parseErr)
	}
	// logger
	if isDebug {
		fmt.Println("parsed results:")
		for _, conf := range config.Resource {
			fmt.Println(conf)
		}
	}

	write(*config, writeFile)
}

func parseArgs() (string, string, bool) {
	var (
		hclFile   = flag.String("f", "main.tf", "the HCL file name you want to parse")
		writeFile = flag.String("w", "main.md", "the Markdown file name you want to write parsed HCL")
		debug     = flag.Bool("d", false, "whether outputs logs for debugging")
	)
	flag.Parse()

	return *hclFile, *writeFile, *debug
}

func parseHcl(fileName string) (*Config, string) {
	parser := hclparse.NewParser()
	f, parseDiags := parser.ParseHCLFile(fileName)
	if parseDiags.HasErrors() {
		return nil, parseDiags.Error()
	}
	var config Config
	if dec := gohcl.DecodeBody(f.Body, nil, &config); dec.HasErrors() {
		return nil, dec.Error()
	}

	return &config, ""
}

func write(config Config, fileName string) {
	file, writeErr := os.Create(fileName)
	if writeErr != nil {
		log.Fatal(writeErr)
	}
	defer file.Close()

	fmt.Fprintln(file, "# Channels")
	fmt.Fprintln(file)

	for _, resource := range config.Resource {
		fmt.Fprintln(file, "##", resource.Name)
		fmt.Fprintln(file)
		fmt.Fprintln(file, *resource.Topic)
		fmt.Fprintln(file)
	}
}
