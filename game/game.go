package game

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/hamao0820/convexHull-drawer/graham"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

var textArea = ebiten.NewImage(ScreenWidth, 15)

func init() {
	textArea.Fill(color.Black)
}

type Game struct {
	plots      []*Plot
	convexHull []*Plot
}

func NewGame() *Game {
	plots := []*Plot{}
	return &Game{
		plots: plots,
	}
}

func (g *Game) DrawConvexHull(screen *ebiten.Image) {
	if len(g.convexHull) < 2 {
		return
	}
	for i := 0; i < len(g.convexHull)-1; i++ {
		vector.StrokeLine(
			screen,
			float32(g.convexHull[i].X()+ScreenWidth/2),
			float32(g.convexHull[i].Y()+ScreenHeight/2),
			float32(g.convexHull[i+1].X()+ScreenWidth/2),
			float32(g.convexHull[i+1].Y()+ScreenHeight/2),
			1,
			color.RGBA{0, 0, 0, 255},
			true,
		)
	}

	vector.StrokeLine(
		screen,
		float32(g.convexHull[len(g.convexHull)-1].X()+ScreenWidth/2),
		float32(g.convexHull[len(g.convexHull)-1].Y()+ScreenHeight/2),
		float32(g.convexHull[0].X()+ScreenWidth/2),
		float32(g.convexHull[0].Y()+ScreenHeight/2),
		1,
		color.RGBA{0, 0, 0, 255},
		true,
	)
}

func (g *Game) Update() error {
	mouseX, mouseY := ebiten.CursorPosition()
	mouseX -= ScreenWidth / 2
	mouseY -= ScreenHeight / 2
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.plots = append(g.plots, NewPlot(mouseX, mouseY))
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		for _, p := range g.plots {
			if p.near(mouseX, mouseY) {
				for i := range g.plots {
					if g.plots[i].id == p.id {
						g.plots = append(g.plots[:i], g.plots[i+1:]...)
						break
					}
				}
				break
			}
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		// Reset
		g.plots = []*Plot{}
		g.convexHull = []*Plot{}
		plotID = 0
	}

	// move
	for i := range g.plots {
		p := g.plots[i]
		p.VX -= float64(p.X()) * 0.0001
		p.VY -= float64(p.Y()) * 0.0001
		p.SetX(p.X() + int(p.VX))
		p.SetY(p.Y() + int(p.VY))
	}

	g.convexHull = graham.Scan(g.plots)

	for i := range g.plots {
		g.plots[i].isConvex = false
	}

	for i := range g.convexHull {
		g.convexHull[i].isConvex = true
	}

	for i := range g.plots {
		if err := g.plots[i].Update(); err != nil {
			return err
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	g.DrawConvexHull(screen)
	for i := range g.plots {
		g.plots[i].Draw(screen)
	}

	screen.DrawImage(textArea, nil)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Num of vertices: %d/%d", len(g.convexHull), len(g.plots)))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func RunGame() {
	g := NewGame()
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Convex Hull Drawer")

	for i := 0; i < 300; i++ {
		p := NewPlot(rand.Intn(ScreenWidth)-ScreenWidth/2, rand.Intn(ScreenHeight)-ScreenHeight/2)
		g.plots = append(g.plots, p)
	}
	g.convexHull = graham.Scan(g.plots)

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
