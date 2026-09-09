// Command tikviewer normalizes a TikTok handle and prints its TikViewer URL.
package main

import (
	"fmt"
	"os"

	"github.com/jacky-xbb/tikviewer-cli/handle"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: tikviewer <handle|profile-url>")
		os.Exit(2)
	}
	h, err := handle.Normalize(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "tikviewer:", err)
		os.Exit(1)
	}
	fmt.Println(handle.StoryURL(h))
}
