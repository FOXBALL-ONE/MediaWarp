package handler

import "testing"

func TestExtractJellyfinItemID(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{name: "standard stream", path: "/Videos/abc123/stream", want: "abc123"},
		{name: "with prefix", path: "/jellyfin/Videos/abc123/original", want: "abc123"},
		{name: "missing videos segment", path: "/Items/abc123", want: ""},
		{name: "trailing slash", path: "/Videos/abc123/stream/", want: "abc123"},
		{name: "upper case segment", path: "/VIDEOS/abc123/stream", want: "abc123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractJellyfinItemID(tt.path); got != tt.want {
				t.Fatalf("extractJellyfinItemID(%s) = %s, want %s", tt.path, got, tt.want)
			}
		})
	}
}
