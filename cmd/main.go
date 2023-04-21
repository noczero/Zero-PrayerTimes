package main

import (
	"embed"

	"github.com/sirupsen/logrus"
	"time"

	"github.com/noczero/Zero-PrayerTimes/config"
	_ "github.com/noczero/Zero-PrayerTimes/logger"
	"github.com/noczero/Zero-PrayerTimes/services"
	"github.com/noczero/Zero-PrayerTimes/utils"
)

//go:embed static/*.mp3
var adhanSound embed.FS

func runForeground(t *time.Time, prayerTimeService *services.AdhanService, playSoundService *services.SoundPLayerServices) {
	prayerTime, err := prayerTimeService.GetPrayerTimes()
	if err != nil {
		logrus.Error(err)
	}
	utils.PlayAdhan(t, prayerTime, playSoundService)
}

func main() {
	logrus.Info("Start Zero-PrayerTimes Adhan")

	// Invoke Services
	playSoundService := services.NewSoundPlayerService(adhanSound)
	prayerTimeService := services.NewAdhanService(config.NewEnv())

	// Use ticker to running every minutes
	ticker := time.NewTicker(1 * time.Minute)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				// trigger when ticker is up
				logrus.Info("Check Prayer Time")

				// use goroutine to play adhan
				go runForeground(&t, prayerTimeService, playSoundService)
			}
		}
	}()

	// Keep the program running forever
	select {}
}
