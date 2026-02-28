package score

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetPage(t *testing.T) {
	t.Parallel()

	now := time.Now()
	older := now.Add(-time.Hour)

	fiveScores := []ScoreEntry{
		{Initials: "AAA", Score: 100, Date: now},
		{Initials: "BBB", Score: 200, Date: now},
		{Initials: "CCC", Score: 300, Date: now},
		{Initials: "DDD", Score: 50, Date: now},
		{Initials: "EEE", Score: 150, Date: now},
	}

	tests := []struct {
		name         string
		scores       []ScoreEntry
		page         int
		size         int
		wantInitials []string
		wantHasMore  bool
	}{
		{
			name:         "empty scores",
			scores:       []ScoreEntry{},
			page:         0,
			size:         5,
			wantInitials: []string{},
			wantHasMore:  false,
		},
		{
			name:         "sorted by score DESC, first page",
			scores:       fiveScores,
			page:         0,
			size:         3,
			wantInitials: []string{"CCC", "BBB", "EEE"},
			wantHasMore:  true,
		},
		{
			name:         "last page has remainder",
			scores:       fiveScores,
			page:         1,
			size:         3,
			wantInitials: []string{"AAA", "DDD"},
			wantHasMore:  false,
		},
		{
			name: "tie broken by date ASC (older first)",
			scores: []ScoreEntry{
				{Initials: "NEW", Score: 100, Date: now},
				{Initials: "OLD", Score: 100, Date: older},
			},
			page:         0,
			size:         2,
			wantInitials: []string{"OLD", "NEW"},
			wantHasMore:  false,
		},
		{
			name:         "exact fit returns false hasMore",
			scores:       fiveScores,
			page:         0,
			size:         5,
			wantInitials: []string{"CCC", "BBB", "EEE", "AAA", "DDD"},
			wantHasMore:  false,
		},
		{
			name:         "page beyond end",
			scores:       []ScoreEntry{{Initials: "AAA", Score: 100, Date: now}},
			page:         1,
			size:         5,
			wantInitials: []string{},
			wantHasMore:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			sm := &ScoreManager{scores: tt.scores}
			got, hasMore := sm.GetPage(tt.page, tt.size)

			initials := make([]string, len(got))
			for i, e := range got {
				initials[i] = e.Initials
			}

			assert.Equal(t, tt.wantInitials, initials)
			assert.Equal(t, tt.wantHasMore, hasMore)
		})
	}
}

func TestSaveAndLoadScores(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "scores.json")

	sm := newScoreManagerAt(path)
	sm.SaveScore("XYZ", 9999, 5, 42)

	sm2 := newScoreManagerAt(path)
	entries, _ := sm2.GetPage(0, 10)

	assert.Len(t, entries, 1)
	assert.Equal(t, "XYZ", entries[0].Initials)
	assert.Equal(t, 9999, entries[0].Score)
	assert.Equal(t, 5, entries[0].Level)
	assert.Equal(t, 42, entries[0].Lines)
}
