package game2048

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
	"math/rand"
	"sort"
	"strconv"
)

// tileSmallFont用于value为个位数的tiles上=====2，4，8
// tileNormalFont用于value为两位数的tiles上====16，32，64
// tileBigFont用于value为三位数的tiles上=======128，256，，，
var (
	tileSmallFont  font.Face
	tileNormalFont font.Face
	tileBigFont    font.Face
)

func init() {
	//创建字体
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatalf("Parse: %v", err)
	}
	//分别配置三种字体
	const dpi = 60
	tileSmallFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatalf("NewFace: %v", err)
	}
	tileNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    32,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatalf("NewFace: %v", err)
	}
	tileBigFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatalf("NewFace: %v", err)
	}
}

const (
	maxMovingCount  = 5
	maxPoppingCount = 6
)

// TileData represents a tile information like a value and a position.
type TileData struct {
	value int
	x     int
	y     int
}

type TilePop struct {
	movingCount       int
	startPoppingCount int
	poppingCount      int
}

// Tile represents a tile information including TileData and animation states.
type Tile struct {
	current TileData
	next    TileData
	pop     TilePop
}

// NewTile creates a new Tile object.
func NewTile(value int, x, y int) *Tile {
	tile := &Tile{
		current: TileData{
			value: value,
			x:     x,
			y:     y,
		},
		pop: TilePop{
			startPoppingCount: maxPoppingCount,
		},
	}
	return tile
}

// IsMoving returns a boolean value indicating if the tile is animating.
func (t *Tile) IsMoving() bool {
	if t.pop.movingCount == 0 {
		return false
	} else {
		return true
	}
}

func (t *Tile) stopAnimation() {
	if t.pop.movingCount != 0 {
		t.current = t.next
		t.next = TileData{}
	}
	t.pop.movingCount = 0
	t.pop.startPoppingCount = 0
	t.pop.poppingCount = 0
}

func tileAt(tiles map[*Tile]struct{}, x, y int) *Tile {
	var result *Tile
	for t := range tiles {
		if t.current.x != x || t.current.y != y {
			continue
		}
		if result != nil {
			panic("not reach")
		}
		result = t
	}
	return result
}

func currentOrNextTileAt(tiles map[*Tile]struct{}, x, y int) *Tile {
	var result *Tile
	for t := range tiles {
		if 0 < t.pop.movingCount {
			if t.next.x != x || t.next.y != y || t.next.value == 0 {
				continue
			}
		} else {
			if t.current.x != x || t.current.y != y {
				continue
			}
		}
		if result != nil {
			panic("not reach")
		}
		result = t
	}
	return result
}

// MoveTiles moves tiles in the given tiles map if possible.
// MoveTiles returns true if there are tiles that are to move, otherwise false.
//
// When MoveTiles is called, all tiles must not be about to move.
func MoveTiles(tiles map[*Tile]struct{}, size int, dir Dir) bool {
	vx, vy := dir.Vector()
	tx := []int{}
	ty := []int{}
	for i := 0; i < size; i++ {
		tx = append(tx, i)
		ty = append(ty, i)
	}
	if vx > 0 {
		sort.Sort(sort.Reverse(sort.IntSlice(tx)))
	}
	if vy > 0 {
		sort.Sort(sort.Reverse(sort.IntSlice(ty)))
	}

	moved := false
	for _, j := range ty {
		for _, i := range tx {
			t := tileAt(tiles, i, j)
			if t == nil {
				continue
			}
			if t.next != (TileData{}) {
				panic("not reach")
			}
			if t.IsMoving() {
				panic("not reach")
			}
			// (ii, jj) is the next position for tile t.
			// (ii, jj) is updated until a mergeable tile is found or
			// the tile t can't be moved any more.
			ii := i
			jj := j
			for {
				ni := ii + vx
				nj := jj + vy
				if ni < 0 || ni >= size || nj < 0 || nj >= size {
					break
				}
				tt := currentOrNextTileAt(tiles, ni, nj)
				if tt == nil {
					ii = ni
					jj = nj
					moved = true
					continue
				}
				if t.current.value != tt.current.value {
					break
				}
				if 0 < tt.pop.movingCount && tt.current.value != tt.next.value {
					// tt is already being merged with another tile.
					// Break here without updating (ii, jj).
					break
				}
				ii = ni
				jj = nj
				moved = true
				break
			}
			// next is the next state of the tile t.
			next := TileData{}
			next.value = t.current.value
			// If there is a tile at the next position (ii, jj), this should be
			// mergeable. Let's merge.
			if tt := currentOrNextTileAt(tiles, ii, jj); tt != t && tt != nil {
				next.value = t.current.value + tt.current.value
				tt.next.value = 0
				tt.next.x = ii
				tt.next.y = jj
				tt.pop.movingCount = maxMovingCount
			}
			next.x = ii
			next.y = jj
			if t.current != next {
				t.next = next
				t.pop.movingCount = maxMovingCount
			}
		}
	}
	if !moved {
		for t := range tiles {
			t.next = TileData{}
			t.pop.movingCount = 0
		}
	}
	return moved
}

func addRandomTile(tiles map[*Tile]struct{}, size int) error {
	cells := make([]bool, size*size)
	for t := range tiles {
		if t.IsMoving() {
			panic("not reach")
		}
		i := t.current.x + t.current.y*size
		cells[i] = true
	}
	availableCells := []int{}
	for i, b := range cells {
		if b {
			continue
		}
		availableCells = append(availableCells, i)
	}
	if len(availableCells) == 0 {
		return errors.New("2048: there is no space to add a new tile")
	}
	c := availableCells[rand.Intn(len(availableCells))]
	v := 2
	if rand.Intn(10) == 0 {
		v = 4
	}
	x := c % size
	y := c / size
	t := NewTile(v, x, y)
	tiles[t] = struct{}{}
	return nil
}

// Update updates the tile's animation states.
func (t *Tile) Update() error {
	switch {
	case 0 < t.pop.movingCount:
		t.pop.movingCount--
		if t.pop.movingCount == 0 {
			if t.current.value != t.next.value && 0 < t.next.value {
				t.pop.poppingCount = maxPoppingCount
			}
			t.current = t.next
			t.next = TileData{}
		}
	case 0 < t.pop.startPoppingCount:
		t.pop.startPoppingCount--
	case 0 < t.pop.poppingCount:
		t.pop.poppingCount--
	}
	return nil
}

func mean(a, b int, rate float64) int {
	return int(float64(a)*(1-rate) + float64(b)*rate)
}

func meanF(a, b float64, rate float64) float64 {
	return a*(1-rate) + b*rate
}

const (
	tileSize   = 80
	tileMargin = 4
)

var (
	tileImage = ebiten.NewImage(tileSize, tileSize)
)

func init() {
	tileImage.Fill(color.White)
}

// Draw draws the current tile to the given boardImage.
func (t *Tile) Draw(boardImage *ebiten.Image) {
	i, j := t.current.x, t.current.y
	ni, nj := t.next.x, t.next.y
	v := t.current.value
	if v == 0 {
		return
	}
	op := &ebiten.DrawImageOptions{}
	x := i*tileSize + (i+1)*tileMargin
	y := j*tileSize + (j+1)*tileMargin
	nx := ni*tileSize + (ni+1)*tileMargin
	ny := nj*tileSize + (nj+1)*tileMargin
	switch {
	case 0 < t.pop.movingCount:
		rate := 1 - float64(t.pop.movingCount)/maxMovingCount
		x = mean(x, nx, rate)
		y = mean(y, ny, rate)
	case 0 < t.pop.startPoppingCount:
		rate := 1 - float64(t.pop.startPoppingCount)/float64(maxPoppingCount)
		scale := meanF(0.0, 1.0, rate)
		op.GeoM.Translate(float64(-tileSize/2), float64(-tileSize/2))
		op.GeoM.Scale(scale, scale)
		op.GeoM.Translate(float64(tileSize/2), float64(tileSize/2))
	case 0 < t.pop.poppingCount:
		const maxScale = 1.2
		rate := 0.0
		if maxPoppingCount*2/3 <= t.pop.poppingCount {
			// 0 to 1
			rate = 1 - float64(t.pop.poppingCount-2*maxPoppingCount/3)/float64(maxPoppingCount/3)
		} else {
			// 1 to 0
			rate = float64(t.pop.poppingCount) / float64(maxPoppingCount*2/3)
		}
		scale := meanF(1.0, maxScale, rate)
		op.GeoM.Translate(float64(-tileSize/2), float64(-tileSize/2))
		op.GeoM.Scale(scale, scale)
		op.GeoM.Translate(float64(tileSize/2), float64(tileSize/2))
	}
	op.GeoM.Translate(float64(x), float64(y))
	op.ColorScale.ScaleWithColor(tileBackgroundColor(v))
	boardImage.DrawImage(tileImage, op)
	str := strconv.Itoa(v)

	f := tileBigFont
	switch {
	case 3 < len(str):
		f = tileSmallFont
	case 2 < len(str):
		f = tileNormalFont
	}

	w := font.MeasureString(f, str).Floor()
	h := (f.Metrics().Ascent + f.Metrics().Descent).Floor()
	x += (tileSize - w) / 2
	y += (tileSize-h)/2 + f.Metrics().Ascent.Floor()
	text.Draw(boardImage, str, f, x, y, tileColor())
}
