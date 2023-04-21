package services

import (
	"github.com/mnadev/adhango/pkg/calc"
	"github.com/mnadev/adhango/pkg/data"
	"github.com/mnadev/adhango/pkg/util"
	"github.com/noczero/Zero-PrayerTimes/config"
	_ "github.com/noczero/Zero-PrayerTimes/logger"
	"time"
)

type AdhanService struct {
	Env *config.Env
}

func NewAdhanService(env *config.Env) *AdhanService {
	return &AdhanService{Env: env}
}

func (service *AdhanService) GetPrayerTimes() (*calc.PrayerTimes, error) {
	date := data.NewDateComponents(time.Now())
	params := calc.GetMethodParameters(calc.SINGAPORE)
	params.Madhab = calc.SHAFI_HANBALI_MALIKI

	coords, err := util.NewCoordinates(service.Env.CityLatitude, service.Env.CityLongitude)
	if err != nil {
		return nil, err
	}

	prayerTimes, err := calc.NewPrayerTimes(coords, date, params)
	if err != nil {
		return nil, err
	}

	err = prayerTimes.SetTimeZone("Asia/Jakarta")
	if err != nil {
		return nil, err
	}

	return prayerTimes, nil
}
