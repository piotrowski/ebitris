package score

import (
	"encoding/json"
	"os"
	"path"
	"slices"
	"time"
)

const saveFile = ".ebitris/scores.json"

type Getter interface {
	GetPage(page, size int) ([]ScoreEntry, bool)
}

type Saver interface {
	SaveScore(initials string, score, level, lines int)
}

type ScoreEntry struct {
	Initials string
	Score    int
	Level    int
	Lines    int
	Date     time.Time
}

type ScoreManager struct {
	scores []ScoreEntry
}

func NewScoreManager() *ScoreManager {
	manager := &ScoreManager{
		scores: []ScoreEntry{},
	}

	if err := ensureBaseDir(saveFile); err != nil {
		panic(err)
	}

	if err := manager.loadScores(); err != nil {
		panic(err)
	}

	return manager
}

func (sm *ScoreManager) SaveScore(initials string, score, level, lines int) {
	entry := ScoreEntry{
		Initials: initials,
		Score:    score,
		Level:    level,
		Lines:    lines,
		Date:     time.Now(),
	}
	sm.scores = append(sm.scores, entry)

	if err := sm.saveScore(); err != nil {
		panic(err)
	}
}

func (sm *ScoreManager) GetPage(page, size int) ([]ScoreEntry, bool) {
	slices.SortFunc(sm.scores, func(a, b ScoreEntry) int {
		if a.Score != b.Score {
			return b.Score - a.Score
		}
		return a.Date.Compare(b.Date)
	})

	start := page * size
	if start >= len(sm.scores) {
		return []ScoreEntry{}, false
	}

	end := start + size
	if end > len(sm.scores) {
		end = len(sm.scores)
		return sm.scores[start:end], false
	}
	return sm.scores[start:end], true
}

func (sm *ScoreManager) saveScore() error {
	jsonData, err := json.Marshal(sm.scores)
	if err != nil {
		return err
	}

	err = os.WriteFile(saveFile, jsonData, 0o600)
	if err != nil {
		return err
	}

	return nil
}

func (sm *ScoreManager) loadScores() error {
	jsonData, err := os.ReadFile(saveFile)
	if err != nil {
		if os.IsNotExist(err) {
			sm.scores = []ScoreEntry{}
			return nil
		}
		return err
	}

	var scores []ScoreEntry
	err = json.Unmarshal(jsonData, &scores)
	if err != nil {
		return err
	}
	sm.scores = scores
	return nil
}

func ensureBaseDir(fpath string) error {
	baseDir := path.Dir(fpath)
	info, err := os.Stat(baseDir)
	if err == nil && info.IsDir() {
		return nil
	}
	return os.MkdirAll(baseDir, 0o755)
}
