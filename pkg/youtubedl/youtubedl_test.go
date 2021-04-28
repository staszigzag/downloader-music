package youtubedl

import (
	"testing"
)

func TestYoutubedl_ExtractVideoID(t *testing.T) {
	d := NewYoutubedl()
	tests := []struct {
		name     string
		urlVideo string
		want     string
		wantErr  bool
	}{
		{
			name:     "error length id",
			urlVideo: "https://www.youtube.com/watch?v=dQw4w9WgXc",
			want:     "",
			wantErr:  true,
		},
		{
			name:     "error domain link",
			urlVideo: "https://www.vk.com/watch?v=dQw4w9WgXcQ",
			want:     "",
			wantErr:  true,
		},
		{
			name:     "error defect id",
			urlVideo: "https://www.youtube.com/watch?v=dQw4<9WgXcQ",
			want:     "",
			wantErr:  true,
		},
		{
			name:     "url 1",
			urlVideo: "http://youtu.be/dQw4w9WgXcQ",
			want:     "dQw4w9WgXcQ",
			wantErr:  false,
		},
		{
			name:     "url 2",
			urlVideo: "http://www.youtube.com/embed/dQw4w9WgXcQ",
			want:     "dQw4w9WgXcQ",
			wantErr:  false,
		},
		{
			name:     "url 3",
			urlVideo: "http://www.youtube.com/watch?v=dQw4w9WgXcQ",
			want:     "dQw4w9WgXcQ",
			wantErr:  false,
		},
		{
			name:     "url 4",
			urlVideo: "http://www.youtube.com/?v=dQw4w9WgXcQ",
			want:     "dQw4w9WgXcQ",
			wantErr:  false,
		},
		{
			name:     "url 5",
			urlVideo: "http://www.youtube.com/v/dQw4w9WgXcQ",
			want:     "dQw4w9WgXcQ",
			wantErr:  false,
		},
		{
			name:     "url 6",
			urlVideo: "http://www.youtube.com/e/dQw4w9WgXcQ",
			want:     "dQw4w9WgXcQ",
			wantErr:  false,
		},
		{
			name:     "url 7",
			urlVideo: "http://www.youtube.com/user/username#p/u/11/dQw4w9WgXcQ",
			want:     "dQw4w9WgXcQ",
			wantErr:  false,
		},
		{
			name:     "url 8",
			urlVideo: "http://www.youtube.com/sandalsResorts#p/c/54B8C800269D7C1B/0/dQw4w9WgXcQ",
			want:     "dQw4w9WgXcQ",
			wantErr:  false,
		},
		{
			name:     "url 9",
			urlVideo: "http://www.youtube.com/watch?feature=player_embedded&v=dQw4w9WgXcQ",
			want:     "dQw4w9WgXcQ",
			wantErr:  false,
		},
		{
			name:     "url 10",
			urlVideo: "http://youtube.com/?feature=player_embedded&v=dQw4w9WgXcQ",
			want:     "dQw4w9WgXcQ",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := d.ExtractVideoID(tt.urlVideo)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractVideoID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExtractVideoID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
