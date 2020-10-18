package systems

import (
	"math/rand"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/yamazaki-ko/game/utils"
)

const (
	// TileNum : Tile数
	TileNum = 200
	// TilePositionY : 初期位置Y
	TilePositionY = 221
	// CellWidth : 1タイル基準幅ß
	CellWidth = 16
	// CellHeight : 1タイル基準高さ
	CellHeight = 16
)

// FallPoint : 落とし穴の位置
var FallPoint []int

// FallStartPoint : 落とし穴の開始位置
var FallStartPoint []int

// Tile :
type Tile struct {
	ecs.BasicEntity
	// RenderComponent:Entityの見た目に関する情報
	common.RenderComponent
	// SpaceComponent:位置に関する情報
	common.SpaceComponent
}

// TileSystem :
type TileSystem struct {
	world      *ecs.World
	tileEntity []*Tile
	texture    *common.Texture
}

// Remove :
func (*TileSystem) Remove(ecs.BasicEntity) {}

// Update :
func (ts *TileSystem) Update(dt float32) {
}

// New :
func (ts *TileSystem) New(w *ecs.World) {

	ts.world = w
	// 落とし穴作成中の状態を保持（0 => 作成していない、1以上 => 作成中）
	tileMakingState := 0
	// 雲の作成中の状態を保持 (0の場合:作成していない、奇数の場合:{(x+1)/2}番目の雲の前半を作成中、偶数の場合:{x/2}番目の雲の後半を作成中)
	cloudMakingState := 0
	// 雲の高さを保持
	cloudHeight := 0
	// 草の作成中の状態を保持 (0の場合:作成していない、奇数の場合:{(x+1)/2}番目の草の前半を作成中、偶数の場合:{x/2}番目の草の後半を作成中)
	glassMakingState := 0
	// 草の作成中の状態を保持 (0の場合:作成していない、奇数の場合:{(x+1)/2}番目の草の前半を作成中、偶数の場合:{x/2}番目の草の後半を作成中)
	mushroomMakingState := 0

	// ランダムにステージを選ぶ
	var tileFile string
	tmp := rand.Intn(2)
	if tmp == 0 {
		tileFile = "./Tilesets/GroundTiles.png"
	} else {
		tileFile = "./Tilesets/GroundTiles.png"
	}
	decorationFile := "./Tilesets/DecorationTiles.png"
	// スプライトシートの作成
	Spritesheet := common.NewSpritesheetWithBorderFromFile(tileFile, CellWidth, CellHeight, 0, 0)
	Tiles := make([]*Tile, 0)

	// 地表の作成
	for j := 0; j <= TileNum; j++ {
		if j > 10 && j < TileNum-50 {
			if tileMakingState > 1 && tileMakingState < 4 {
				for t := 0; t < 8; t++ {
					FallPoint = append(FallPoint, j*CellWidth-t)
				}
			} else if tileMakingState == 0 {
				// たまに落とし穴を作る
				randomNum := rand.Intn(10)
				if randomNum == 0 {
					FallStartPoint = append(FallStartPoint, j*CellWidth)
					tileMakingState = 1
				}
			}
		}
		// 描画するタイルを保持
		var selectedTile int
		// 描画するタイルを選択
		switch tileMakingState {
		case 0:
			selectedTile = 1
		case 1:
			selectedTile = 5
		case 2:
			tileMakingState++
			continue
		case 3:
			tileMakingState++
			continue
		case 4:
			selectedTile = 0
		}
		// タイルEntityの作成
		tile := &Tile{BasicEntity: ecs.NewBasic()}
		// 位置情報の設定
		tile.SpaceComponent.Position = engo.Point{
			X: float32(j * CellWidth),
			Y: float32(TilePositionY + CellHeight),
		}
		// 見た目の設定
		tile.RenderComponent.Drawable = Spritesheet.Cell(selectedTile)
		tile.RenderComponent.SetZIndex(0)
		Tiles = append(Tiles, tile)

		if tileMakingState > 0 {
			if tileMakingState == 4 {
				tileMakingState = 0
				continue
			}
			tileMakingState++
		}
	}
	// 地面の描画
	for i := 0; i < 3; i++ {
		tileMakingState = 0
		for j := 0; j <= TileNum; j++ {
			if tileMakingState == 0 {
				// 落とし穴を作る場合
				if utils.Contains(FallStartPoint, j*CellWidth) {
					tileMakingState = 1
				}
			}
			// 描画するタイルを保持
			var selectedTile int
			// 描画するタイルを選択
			switch tileMakingState {
			case 0:
				selectedTile = 31
			case 1:
				selectedTile = 35
			case 2:
				tileMakingState++
				continue
			case 3:
				tileMakingState++
				continue
			case 4:
				selectedTile = 30
			}
			tile := &Tile{BasicEntity: ecs.NewBasic()}
			tile.SpaceComponent.Position = engo.Point{
				X: float32(j * CellWidth),
				Y: float32(285 - i*CellHeight),
			}
			tile.RenderComponent.Drawable = Spritesheet.Cell(selectedTile)
			tile.RenderComponent.SetZIndex(0)
			Tiles = append(Tiles, tile)

			if tileMakingState > 0 {
				if tileMakingState == 4 {
					tileMakingState = 0
					continue
				}
				tileMakingState++
			}
		}
	}

	// スプライトシートの再作成
	Spritesheet = common.NewSpritesheetWithBorderFromFile(decorationFile, CellWidth, CellHeight, 0, 0)
	for j := 0; j <= TileNum; j++ {
		// 雲の作成
		if cloudMakingState == 0 {
			randomNum := rand.Intn(12)
			if randomNum > 5 && randomNum < 12 && randomNum%2 == 0 {
				cloudMakingState = randomNum
			}
			cloudHeight = rand.Intn(70) + 10
		}
		if cloudMakingState != 0 {
			// 雲Entityの作成
			cloudTile := cloudMakingState
			cloud := &Tile{BasicEntity: ecs.NewBasic()}
			cloud.SpaceComponent.Position = engo.Point{
				X: float32(j * CellWidth),
				Y: float32(cloudHeight),
			}
			cloud.RenderComponent.Drawable = Spritesheet.Cell(cloudTile)
			cloud.RenderComponent.SetZIndex(0)
			Tiles = append(Tiles, cloud)
			// 前半を作成中であれば、次は後半を作成する
			if cloudMakingState%2 == 0 {
				cloudMakingState++
			} else {
				cloudMakingState = 0
			}
		}
		// 草、キノコの作成
		// 落とし穴の上には作らない
		if !utils.Contains(FallPoint, j*CellWidth) {
			var grassTile int
			var mushroomTile int
			// 草の作成
			if glassMakingState == 0 {
				randomNum := rand.Intn(18)
				if randomNum < 4 && randomNum%2 == 0 {
					glassMakingState = 18 + randomNum
				}
			}
			if glassMakingState != 0 {
				if glassMakingState%2 == 0 && utils.Contains(FallPoint, (j+1)*CellWidth) {

				} else {
					grassTile = glassMakingState
					grass := &Tile{BasicEntity: ecs.NewBasic()}
					grass.SpaceComponent.Position = engo.Point{
						X: float32(j * CellWidth),
						Y: float32(TilePositionY),
					}
					grass.RenderComponent.Drawable = Spritesheet.Cell(grassTile)
					grass.RenderComponent.SetZIndex(1)
					Tiles = append(Tiles, grass)
					// 前半を作成中であれば、次は後半を作成する
					if glassMakingState%2 == 0 {
						glassMakingState++
					} else {
						glassMakingState = 0
					}
				}
			}
			// キノコの作成
			if mushroomMakingState == 0 {
				randomNum := rand.Intn(18)
				if randomNum > 10 && randomNum < 17 && randomNum%2 == 0 {
					mushroomMakingState = 18 + randomNum
				}
			}
			if mushroomMakingState != 0 {
				if mushroomMakingState%2 == 0 && utils.Contains(FallPoint, (j+1)*CellWidth) {

				} else {
					mushroomTile = mushroomMakingState
					mushroom := &Tile{BasicEntity: ecs.NewBasic()}
					mushroom.SpaceComponent.Position = engo.Point{
						X: float32(j * CellWidth),
						Y: float32(TilePositionY),
					}
					mushroom.RenderComponent.Drawable = Spritesheet.Cell(mushroomTile)
					mushroom.RenderComponent.SetZIndex(1)
					Tiles = append(Tiles, mushroom)
					// 前半を作成中であれば、次は後半を作成する
					if mushroomMakingState%2 == 0 {
						mushroomMakingState++
					} else {
						mushroomMakingState = 0
					}
				}
			}
		}

	}

	tileMakingState = 0
	for _, system := range ts.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			for _, v := range Tiles {
				ts.tileEntity = append(ts.tileEntity, v)
				sys.Add(&v.BasicEntity, &v.RenderComponent, &v.SpaceComponent)
			}
		}
	}
}
