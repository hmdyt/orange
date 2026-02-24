package client

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hmdyt/orange/adapter/client/api"
)

//go:embed fonts/MPLUS1p-Regular.ttf
var fontData []byte

const (
	screenWidth  = 640
	screenHeight = 480
)

type loginResult struct {
	message string
	userID  string
	err     error
}

type Game struct {
	apiClient *api.Client
	resultCh  chan loginResult
	message   string
	userID    string
	errMsg    string
	fontFace  *text.GoTextFace
}

func NewGame(apiClient *api.Client) *Game {
	source, err := text.NewGoTextFaceSource(bytes.NewReader(fontData))
	if err != nil {
		log.Fatalf("failed to load font: %v", err)
	}
	g := &Game{
		apiClient: apiClient,
		resultCh:  make(chan loginResult, 1),
		fontFace: &text.GoTextFace{
			Source: source,
			Size:   24,
		},
	}
	go g.login()
	return g
}

func (g *Game) login() {
	resp, err := g.apiClient.Login(context.Background(), "プレイヤー1")
	if err != nil {
		g.resultCh <- loginResult{err: err}
		return
	}
	g.resultCh <- loginResult{
		message: resp.Message,
		userID:  resp.UserId,
	}
}

func (g *Game) Update() error {
	select {
	case result := <-g.resultCh:
		if result.err != nil {
			g.errMsg = fmt.Sprintf("Error: %v", result.err)
		} else {
			g.message = result.message
			g.userID = result.userID
		}
	default:
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{30, 30, 50, 255})

	if g.errMsg != "" {
		g.drawCenteredText(screen, g.errMsg, screenHeight/2)
		return
	}
	if g.message == "" {
		g.drawCenteredText(screen, "ログイン中...", screenHeight/2)
		return
	}

	g.drawCenteredText(screen, g.message, screenHeight/2-30)
	g.drawCenteredText(screen, fmt.Sprintf("User ID: %s", g.userID), screenHeight/2+20)
}

func (g *Game) drawCenteredText(screen *ebiten.Image, str string, y float64) {
	w, _ := text.Measure(str, g.fontFace, 0)
	op := &text.DrawOptions{}
	op.GeoM.Translate((screenWidth-w)/2, y)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, str, g.fontFace, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
