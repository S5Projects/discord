package discord_test

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/S5Projects/discord"
)

var webhookURL = ""

func TestMain(m *testing.M) {
	for _, env := range os.Environ() {
		components := strings.Split(env, "=")
		key := components[0]
		value := components[1]
		if key == "DISCORD_WEBHOOK_URL" {
			webhookURL = value
			break
		}
	}

	if webhookURL == "" {
		fmt.Fprintf(os.Stderr, "DISCORD_WEBHOOK_URL environment variable is required\n")
		os.Exit(1)
	}

	code := m.Run()
	os.Exit(code)
}

func TestSay(t *testing.T) {
	wh, err := discord.New(webhookURL)
	if err != nil {
		panic(err)
	}
	err = wh.Say("Hello, world!")
	if err != nil {
		t.Errorf("Error posting plain-text message: %s", err.Error())
	}
}

func TestPost(t *testing.T) {
	wh, err := discord.New(webhookURL)
	if err != nil {
		panic(err)
	}
	err = wh.Post(discord.PostOptions{
		Content: "Hello, world!",
		Embeds: []discord.Embed{
			{
				Color: 16777215,
				Author: &discord.Author{
					Name: "ecnepsnai",
					URL:  "https://github.com/ecnepsnai",
				},
				Title:       "Amazing!",
				Description: "This is a cool embed",
			},
		},
	})
	if err != nil {
		t.Errorf("Error posting complex message: %s", err.Error())
	}
}

func TestFileUpload(t *testing.T) {
	wh, err := discord.New(webhookURL)
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(path.Join(".", "discord.go"), os.O_RDONLY, 0644)
	if err != nil {
		t.Fatalf("Error opening file: %s", err.Error())
	}
	defer f.Close()
	content := discord.PostOptions{
		Content: "Hello, world!",
	}
	fileOptions := discord.FileOptions{
		FileName: "discord.go",
		Reader:   f,
	}
	if err := wh.UploadFile(content, fileOptions); err != nil {
		t.Errorf("Error posting message with file attachment: %s", err.Error())
	}
}
