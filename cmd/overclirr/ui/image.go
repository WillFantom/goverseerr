package ui

import (
	"os"
	"runtime"

	"github.com/eliukblau/pixterm/pkg/ansimage"
	colorful "github.com/lucasb-eyer/go-colorful"
	"github.com/pkg/browser"
	"github.com/sirupsen/logrus"
	"github.com/willfantom/goverseerr"
	"golang.org/x/term"
)

func ShowMediaPoster(posterPath string) {
	termErr := printPosterInTerminal(goverseerr.PosterPathBase + posterPath)
	if termErr != nil {
		logrus.WithFields(logrus.Fields{
			"extended":   termErr.Error(),
			"posterPath": posterPath,
		}).Errorln("could not print the poster in the terminal")
	} else {
		return
	}
	err := browser.OpenURL(goverseerr.PosterPathBase + posterPath)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"extended":   err.Error(),
			"posterPath": posterPath,
		}).Fatalln("could not show poster in the browser")
	}
}

func printPosterInTerminal(url string) error {
	var image *ansimage.ANSImage
	var x, y int
	var err error
	if term.IsTerminal(int(os.Stdout.Fd())) {
		x, y, err = term.GetSize(int(os.Stdout.Fd()))
	}
	if err != nil {
		x = 80
		y = 24
	}
	sm := ansimage.ScaleMode(2)
	dm := ansimage.DitheringMode(0)
	sfy, sfx := ansimage.BlockSizeY, ansimage.BlockSizeX
	if ansimage.DitheringMode(0) == ansimage.NoDithering {
		sfy, sfx = 2, 1
	}
	mc, _ := colorful.Hex("#000000")
	image, err = ansimage.NewScaledFromURL(url, sfy*y, sfx*x, mc, sm, dm)
	if err != nil {
		return err
	}
	if term.IsTerminal(int(os.Stdout.Fd())) {
		ansimage.ClearTerminal()
	}
	image.SetMaxProcs(runtime.NumCPU())
	image.DrawExt(false, false)
	return nil
}
