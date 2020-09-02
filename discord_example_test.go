package discord_test

import (
	"github.com/S5Projects/discord"
)

func ExampleSay() {
	wh, err := discord.New("https://discord.com/api/webhooks/.../...")
	if err != nil {
		panic(err)
	}
	wh.Say("Hello, world!")
}

func ExamplePost() {
	wh, err := discord.New("https://discord.com/api/webhooks/.../...")
	if err != nil {
		panic(err)
	}
	wh.Post(discord.PostOptions{
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
}

func ExampleUploadFile() {
	wh, err := discord.New("https://discord.com/api/webhooks/.../...")
	if err != nil {
		panic(err)
	}
	// var f *io.Reader // Pretend we've opened a file
	content := discord.PostOptions{
		Content: "Hello, world!",
	}
	fileOptions := discord.FileOptions{
		FileName: "my_hot_mixtape.mp3",
		// Reader:   f,
	}
	wh.UploadFile(content, fileOptions)
}
