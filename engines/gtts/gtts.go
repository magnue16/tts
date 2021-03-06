package gtts

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"

	"github.com/kladd/tts"
)

// Speaker implements the Speaker interface for Google TTS
type Speaker struct{}

const urlFmt = "http://translate.google.com/translate_tts?ie=UTF-8&tl=%s&q=%s"

// Say says phrase aloud using Google TTS
func (s *Speaker) Say(phrase string) {
	f, err := ioutil.TempFile(os.TempDir(), "gtts")
	defer f.Close()
	defer os.Remove(f.Name())

	r, err := http.Get(
		fmt.Sprintf(
			urlFmt,
			tts.Lang,
			url.QueryEscape(phrase),
		),
	)
	defer r.Body.Close()

	_, err = io.Copy(f, r.Body)
	if err != nil {
		fmt.Println(err)
	}

	exec.Command(tts.PlayCmd, f.Name()).Run()
}
