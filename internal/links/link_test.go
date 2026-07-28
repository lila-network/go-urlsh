package links

import "testing"

func TestLink_IsValid(t *testing.T) {
	tests := []struct {
		name string
		link Link
		want bool
	}{
		{
			name: "valid link",
			link: Link{ID: "test", URL: "https://example.com"},
			want: true,
		},
		{
			name: "empty ID",
			link: Link{ID: "", URL: "https://example.com"},
			want: false,
		},
		{
			name: "empty URL",
			link: Link{ID: "test", URL: ""},
			want: false,
		},
		{
			name: "empty ID and URL",
			link: Link{ID: "", URL: ""},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.link.IsValid(); got != tt.want {
				t.Errorf("Link.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
