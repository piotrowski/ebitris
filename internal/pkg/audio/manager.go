package audio

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

type EffectPlayer interface {
	PlayEffect(name EffectName)
}

type MusicPlayer interface {
	StartPlaylist(names ...SongName)
}

type AudioUpdater interface {
	Update() error
}

const samplingRate = 44100

type AudioManager struct {
	audioContext *audio.Context

	decodedSongs   map[SongName][]byte
	decodedEffects map[EffectName][]byte

	musicPlayer   *audio.Player
	effectPlayers map[EffectName]*audio.Player

	currentSong SongName
	playlist    []SongName
}

func NewAudioManager() *AudioManager {
	audioCtx := audio.NewContext(samplingRate)

	decodedSongs := make(map[SongName][]byte)
	for name, raw := range songs {
		data, err := decodeMP3(raw)
		if err != nil {
			slog.Error("failed to decode audio", "subsystem", "audio", "song", name, "err", err)
			panic(err)
		}
		decodedSongs[name] = data
	}

	decodedEffects := make(map[EffectName][]byte)
	for name, raw := range effects {
		data, err := decodeMP3(raw)
		if err != nil {
			slog.Error("failed to decode audio", "subsystem", "audio", "effect", name, "err", err)
			panic(err)
		}
		decodedEffects[name] = data
	}

	// For better effect use 4 players for each effect and overlap them.
	effectPlayers := make(map[EffectName]*audio.Player)
	for name, data := range decodedEffects {
		p, err := audioCtx.NewPlayer(bytes.NewReader(data))
		if err != nil {
			slog.Error("failed to create effect player", "subsystem", "audio", "effect", name, "err", err)
			panic(err)
		}
		effectPlayers[name] = p
	}

	return &AudioManager{
		audioContext:   audioCtx,
		decodedSongs:   decodedSongs,
		decodedEffects: decodedEffects,

		effectPlayers: effectPlayers,
	}
}

func (m *AudioManager) Update() error {
	if m.musicPlayer != nil && !m.musicPlayer.IsPlaying() && len(m.playlist) > 1 {
		m.PlayShuffle()
	}

	return nil
}

func (m *AudioManager) PlayShuffle() {
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

func (m *AudioManager) PlaySong(name SongName) {
	slog.Info("playing track", "subsystem", "audio", "song", name)
	m.currentSong = name

	if m.musicPlayer != nil {
		m.musicPlayer.Pause()
	}

	if err := m.loadTrack(name); err != nil {
		slog.Error("failed to load track", "subsystem", "audio", "song", name, "err", err)
		panic(err)
	}
	if err := m.musicPlayer.Rewind(); err != nil {
		slog.Error("failed to rewind track", "subsystem", "audio", "song", name, "err", err)
		panic(err)
	}
	m.musicPlayer.SetVolume(0.1)
	m.musicPlayer.Play()
}

func (m *AudioManager) StartPlaylist(names ...SongName) {
	m.playlist = names
	m.PlayShuffle()
}

func (m *AudioManager) loadTrack(name SongName) error {
	var err error
	if m.musicPlayer != nil {
		err = m.musicPlayer.Close()
		if err != nil {
			return err
		}
	}

	m.musicPlayer, err = m.audioContext.NewPlayer(bytes.NewReader(m.decodedSongs[name]))
	if err != nil {
		return err
	}

	return nil
}

func (m *AudioManager) PlayEffect(name EffectName) {
	p, ok := m.effectPlayers[name]
	if !ok {
		return
	}
	if err := p.Rewind(); err != nil {
		slog.Error("failed to rewind track", "subsystem", "audio", "effect", name, "err", err)
	}

	p.Play()
}

func decodeMP3(raw []byte) ([]byte, error) {
	stream, err := mp3.DecodeWithSampleRate(samplingRate, bytes.NewReader(raw))
	if err != nil {
		return nil, fmt.Errorf("failed to decode mp3 file: %w", err)
	}
	data, err := io.ReadAll(stream)
	if err != nil {
		return nil, fmt.Errorf("failed to read mp3 file: %w", err)
	}
	return data, nil
}
