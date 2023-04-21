package utils

import (
	"github.com/mnadev/adhango/pkg/calc"
	_ "github.com/noczero/Zero-PrayerTimes/logger"
	"github.com/noczero/Zero-PrayerTimes/services"
	"github.com/sirupsen/logrus"
	"time"
)

func compareTwoTimesInMinutes(t1 *time.Time, t2 time.Time) bool {
	if t1.Hour() == t2.Hour() && t1.Minute() == t2.Minute() {
		return true
	}
	return false
}

func PlayAdhan(currentTime *time.Time, prayerTime *calc.PrayerTimes, soundPlayer *services.SoundPLayerServices) bool {

	listPrayerTimes := map[string]time.Time{
		"MAGHRIB": prayerTime.Maghrib,
		"ISHA":    prayerTime.Isha,
		"FAJR":    prayerTime.Fajr,
		"DHUHR":   prayerTime.Dhuhr,
		"ASR":     prayerTime.Asr,
	}

	for key, value := range listPrayerTimes {
		if compareTwoTimesInMinutes(currentTime, value) {
			logrus.Info("Current Prayer : ", key)

			if key == "FAJR" {
				_, err := soundPlayer.PlaySound(true)
				if err != nil {
					logrus.Error(err)
				}
			} else {
				_, err := soundPlayer.PlaySound(false)
				if err != nil {
					logrus.Error(err)
				}
			}

			return true
		} else {
			logrus.Debug("Skipped - ", key, value)
		}
	}

	return false
}
