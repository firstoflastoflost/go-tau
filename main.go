package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const SettingsFileName = "settings.json"

type App struct {
	Urls    []Url   `json:"urls"`
	Options Options `json:"options"`
}

type Url struct {
	Address string `json:"address"`
	Status  Status
}

type Options struct {
	SslMode string        `json:"ssl_mode"`
	Timeout time.Duration `json:"timeout"`
}

type Status struct {
	Success   bool
	HttpCode  int
	Timestamp time.Time
}

func main() {
	app := parseConfig()

	checkUrls(&app)

	file, _ := json.MarshalIndent(app.Urls, "", "")

	resultFileName := "report_" + time.Now().Format("2006-01-02 15:04:05") + ".json"

	_ = os.WriteFile(resultFileName, file, 0644)
}

func parseConfig() App {
	jsonFile, err := os.Open(SettingsFileName)

	if err != nil {
		fmt.Println("Error opened file urls.json")
	}

	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println("Error closed file urls.json")
			return
		}
	}(jsonFile)

	byteValue, _ := io.ReadAll(jsonFile)

	var app App

	json.Unmarshal(byteValue, &app)

	for i, url := range app.Urls {
		app.Urls[i].Address = app.Options.SslMode + "://" + url.Address
	}

	return app
}

func checkUrls(app *App) {

	timeout := app.Options.Timeout

	for i := range app.Urls {
		app.Urls[i].Status = probeUrl(app.Urls[i].Address, timeout)
	}
}

func probeUrl(address string, timeout time.Duration) Status {
	client := http.Client{
		Timeout: timeout * time.Second,
	}

	resp, err := client.Get(address)
	if err != nil {
		return Status{
			Success:   false,
			HttpCode:  0,
			Timestamp: time.Now(),
		}
	}
	defer resp.Body.Close()

	return Status{
		Success:   true,
		HttpCode:  resp.StatusCode,
		Timestamp: time.Now(),
	}
}
