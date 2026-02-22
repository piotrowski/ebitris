package music

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

type AudioPlayer struct {
	audioContext *audio.Context

	musicPlayer   *audio.Player
	effectsPlayer *audio.Player

	currentSong SongName
	playlist    []SongName

	autoPlay      chan bool
	currentEffect EffectName
}

const samplingRate = 44100

func NewManager() *AudioPlayer {
	return &AudioPlayer{
		autoPlay:     nil,
		audioContext: audio.NewContext(samplingRate),
	}
}

func (m *AudioPlayer) StartAutoPlay() {
	m.autoPlay = make(chan bool, 1)

	if len(m.playlist) == 0 {
		return
	}

	for {
		select {
		case <-m.autoPlay:
			return
		default:
			if m.musicPlayer == nil {
				close(m.autoPlay)
				return
			}

			if m.musicPlayer.IsPlaying() {
				time.Sleep(100 * time.Millisecond)
				continue
			}
			m.PlayShuffle()
		}
	}
}

func (m *AudioPlayer) StopAutoPlay() {
	close(m.autoPlay)
	if m.musicPlayer != nil {
		m.musicPlayer.Pause()
	}
}

func (m *AudioPlayer) SetPlaylist(names ...SongName) {
	m.playlist = names
}

func (m *AudioPlayer) PlayShuffle() {
	if len(m.playlist) == 0 {
		return
	}

	var randomIndex int
	for i := 0; i < 10; i++ {
		randomIndex = rand.Intn(len(m.playlist))
		if m.playlist[randomIndex] != m.currentSong {
			break
		}
	}

	m.PlaySong(m.playlist[randomIndex])
}

func (m *AudioPlayer) PlaySong(name SongName) {
	if m.currentSong == name {
		return
	}

	m.currentSong = name

	if m.musicPlayer != nil {
		m.musicPlayer.Pause()
	}

	if err := m.loadTrack(name); err != nil {
		panic(err)
	}
	if err := m.musicPlayer.Rewind(); err != nil {
		panic(err)
	}
	m.musicPlayer.Play()
	m.musicPlayer.SetVolume(0)
	go func() {
		player := m.musicPlayer
		for i := 0; i <= 10; i++ {
			if player != m.musicPlayer {
				return
			}

			time.Sleep(100 * time.Millisecond)
			fmt.Println("Volume: ", float64(i)/10)
			player.SetVolume(float64(i) / 10)
		}
	}()

	fmt.Printf("Playing song: %v\n", name)
}

func (m *AudioPlayer) loadTrack(name SongName) error {
	src := bytes.NewReader(songs[name])
	stream, err := mp3.DecodeWithSampleRate(samplingRate, src)
	if err != nil {
		return err
	}

	if m.musicPlayer != nil {
		err = m.musicPlayer.Close()
		if err != nil {
			return err
		}
	}

	m.musicPlayer, err = m.audioContext.NewPlayer(stream)
	if err != nil {
		return err
	}

	return nil
}

func (m *AudioPlayer) PlayEffect(name EffectName) {
	if err := m.loadEffect(name); err != nil {
		panic(err)
	}
	err := m.effectsPlayer.Rewind()
	if err != nil {
		panic(err)
	}
	m.effectsPlayer.Play()
}

func (m *AudioPlayer) loadEffect(name EffectName) error {
	m.currentEffect = name
	src := bytes.NewReader(effects[name])
	stream, err := mp3.DecodeWithSampleRate(samplingRate, src)
	if err != nil {
		return err
	}

	if m.effectsPlayer != nil {
		go func() {
			player := m.effectsPlayer
			for player.IsPlaying() {
				time.Sleep(50 * time.Millisecond)
			}
			err = player.Close()
			if err != nil {
				fmt.Printf("Error closing effect player: %v\n", err)
			}
		}()
	}
	m.effectsPlayer, err = m.audioContext.NewPlayer(stream)
	if err != nil {
		return err
	}

	return nil
}
