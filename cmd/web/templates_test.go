package main

import (
	"testing"
	"time"

	"github.com/PrayasPathak/snippetbox/internal/assert"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2024, 8, 16, 10, 34, 0, 0, time.UTC),
			want: "16 Aug 2024 at 10:34"},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm: time.Date(2024, 8, 20, 10, 15, 0, 0, time.FixedZone("CET",
				1*60*60)),
			want: "20 Aug 2024 at 10:15",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			hd := humanDate(test.tm)
			assert.Equal(t, hd, test.want)
		})
	}
}
