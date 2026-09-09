package handle

import "testing"

func TestNormalize(t *testing.T) {
	ok := map[string]string{
		"@Charli":                          "charli",
		"charli":                           "charli",
		"https://www.tiktok.com/@Charli":   "charli",
		"tiktok.com/@charli/video/12345":   "charli",
		"https://tiktok.com/@charli?lang=en": "charli",
	}
	for in, want := range ok {
		got, err := Normalize(in)
		if err != nil || got != want {
			t.Errorf("Normalize(%q) = %q, %v; want %q", in, got, err, want)
		}
	}
	for _, in := range []string{"", "@", "a", "has space", "tiktok.com/"} {
		if _, err := Normalize(in); err == nil {
			t.Errorf("Normalize(%q) should fail", in)
		}
	}
}

func TestStoryURL(t *testing.T) {
	if got := StoryURL("charli"); got != "https://tikviewer.org/@charli" {
		t.Errorf("StoryURL = %q", got)
	}
}
