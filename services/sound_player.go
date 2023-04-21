package services

import (
	"bytes"
	"embed"
	"errors"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
	_ "github.com/noczero/Zero-PrayerTimes/logger"
	"github.com/sirupsen/logrus"
	"time"
)

type SoundPLayerServices struct {
	File embed.FS
}

func NewSoundPlayerService(file embed.FS) *SoundPLayerServices {
	return &SoundPLayerServices{File: file}
}

func (service *SoundPLayerServices) PlaySound(isFajr bool) (bool, error) {

	// read file
	fileBytes, err := service.File.ReadFile("static/" + "adhan-with-doa.mp3")

	if isFajr {
		fileBytes, err = service.File.ReadFile("static/" + "adhan-fajr.mp3")
	}

	if err != nil {
		return false, errors.New("reading my-file.mp3 failed: " + err.Error())
	}

	// Convert the pure bytes into a reader object that can be used with the mp3 decoder
	fileBytesReader := bytes.NewReader(fileBytes)

	// Decode file
	decodedMp3, err := mp3.NewDecoder(fileBytesReader)
	if err != nil {
		return false, errors.New("mp3.NewDecoder failed: " + err.Error())
	}

	// Prepare an Oto context (this will use your default audio device) that will
	// play all our sounds. Its configuration can't be changed later.

	// Usually 44100 or 48000. Other values might cause distortions in Oto
	samplingRate := 44100

	// Number of channels (aka locations) to play sounds from. Either 1 or 2.
	// 1 is mono sound, and 2 is stereo (most speakers are stereo).
	numOfChannels := 2

	// Bytes used by a channel to represent one sample. Either 1 or 2 (usually 2).
	audioBitDepth := 2

	// Remember that you should **not** create more than one context
	otoCtx, readyChan, err := oto.NewContext(samplingRate, numOfChannels, audioBitDepth)
	if err != nil {
		return false, errors.New("oto.NewContext failed: " + err.Error())
	}
	// It might take a bit for the hardware audio devices to be ready, so we wait on the channel.
	<-readyChan

	// Create a new 'player' that will handle our sound. Paused by default.
	player := otoCtx.NewPlayer(decodedMp3)

	// Play starts playing the sound and returns without waiting for it (Play() is async).
	logrus.Info("Playing Adhan")
	player.Play()

	// We can wait for the sound to finish playing using something like this
	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}

	// If you don't want the player/sound anymore simply close
	err = player.Close()
	if err != nil {
		return false, errors.New("player.Close failed: " + err.Error())
	}

	return true, nil
}
