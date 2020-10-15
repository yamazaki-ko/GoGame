package systems

import (
	"fmt"
	"strconv"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/yamazaki-ko/game/utils"
)

const (
	// MoveDistance : 移動距離
	MoveDistance = 5
	// JumpHeight : ジャンプの高さ
	JumpHeight = 5
	// JumpTime : ジャンプの時間
	JumpTime = 12
	// PlayerPositionY : 初期位置Y
	PlayerPositionY = 210
	// SpriteSheetCell : スプライトシートで使用するセル番号
	SpriteSheetCell = 28
)

// Player :
type Player struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	// スプライトシート
	Spritesheet *common.Spritesheet
	// プレイヤーの状態
	pattern int
	// ジャンプの時間
	jumpDuration int
	// 2段ジャンプの時間
	jumpDuration2Step int
	// ジャンプの回数
	jumpCount int
	// ジャンプ初期位置
	jumpPositionY int
	// カメラの進んだ距離
	distance int
	// 落ちているかどうか
	ifFalling bool
	// ダメージ
	damage int
	// GOALしたかどうか
	ifGoal bool
}

// PlayerSystem :
type PlayerSystem struct {
	world        *ecs.World
	playerEntity *Player
	texture      *common.Texture
}

var playerFile = "./Free Hero/LongHair_BlueTunic_Shield COMPRESSED.png"

// Remove :
func (*PlayerSystem) Remove(ecs.BasicEntity) {}

// Update :
func (ps *PlayerSystem) Update(dt float32) {
	// ダメージが1であればゲームを終了
	if ps.playerEntity.damage > 0 {
		ps.playerEntity.damage = -1
		fmt.Println("死亡位置は" + strconv.Itoa(ps.playerEntity.distance))
		whenDied(ps)
	}
	// 落とし穴に落ちる
	if ps.playerEntity.jumpDuration == 0 && utils.Contains(FallPoint, int(ps.playerEntity.SpaceComponent.Position.X)) {
		ps.playerEntity.ifFalling = true
		ps.playerEntity.SpaceComponent.Position.Y += MoveDistance
	}
	// 穴に落ち切ったらライフを0にする
	if ps.playerEntity.SpaceComponent.Position.Y > 300 {
		if ps.playerEntity.damage == 0 {
			ps.playerEntity.damage++
		}
	}
	// プレーヤーを右に移動
	if !ps.playerEntity.ifFalling {
		if engo.Input.Button("MoveRight").Down() {
			// Goal地点に達したら右移動はしない
			if int(ps.playerEntity.SpaceComponent.Position.X) >= GoalPoint {

			} else {
				// 画面の真ん中より左に位置していれば、カメラを移動せずプレーヤーを移動する
				if int(ps.playerEntity.SpaceComponent.Position.X) < ps.playerEntity.distance+int(engo.WindowWidth())/2 {
					ps.playerEntity.SpaceComponent.Position.X += 5
				} else {
					// 画面の右端に達していなければプレーヤーを移動する
					if int(ps.playerEntity.SpaceComponent.Position.X) < ps.playerEntity.distance+int(engo.WindowWidth())-10 {
						ps.playerEntity.SpaceComponent.Position.X += 5
					}
					// カメラを移動する
					engo.Mailbox.Dispatch(common.CameraMessage{
						Axis:        common.XAxis,
						Value:       5,
						Incremental: true,
					})
					ps.playerEntity.distance += MoveDistance
					utils.SetPosition(ps.playerEntity.distance)
				}
				if ps.playerEntity.jumpDuration == 0 {
					switch ps.playerEntity.pattern {
					case 0:
						ps.playerEntity.pattern = 1
					case 1:
						ps.playerEntity.pattern = 2
					case 2:
						ps.playerEntity.pattern = 3
					case 3:
						ps.playerEntity.pattern = 4
					case 4:
						ps.playerEntity.pattern = 0
					}
					ps.playerEntity.RenderComponent.Drawable = ps.playerEntity.Spritesheet.Cell(SpriteSheetCell + ps.playerEntity.pattern)

				}
			}
		}
	}
	// プレーヤーを左に移動
	if engo.Input.Button("MoveLeft").Down() {
		if int(ps.playerEntity.SpaceComponent.Position.X) > ps.playerEntity.distance+10 {
			ps.playerEntity.SpaceComponent.Position.X -= MoveDistance
		}
	}
	// プレーヤーをジャンプ
	if engo.Input.Button("Jump").JustPressed() {
		// 2段ジャンプ
		if ps.playerEntity.jumpCount == 1 {
			ps.playerEntity.jumpDuration = 1
			ps.playerEntity.jumpDuration2Step = ps.playerEntity.jumpDuration
			ps.playerEntity.jumpCount = 0
		}
		// 初回ジャンプ
		if ps.playerEntity.jumpDuration == 0 {
			ps.playerEntity.jumpDuration = 1
			ps.playerEntity.jumpDuration2Step = JumpTime
			ps.playerEntity.jumpCount = 1
			ps.playerEntity.jumpPositionY = int(ps.playerEntity.SpaceComponent.Position.Y)
		}
	}
	if ps.playerEntity.jumpDuration != 0 {
		ps.playerEntity.jumpDuration++
		if ps.playerEntity.jumpDuration < 2+JumpTime {
			ps.playerEntity.SpaceComponent.Position.Y -= JumpHeight
		} else if ps.playerEntity.jumpDuration < 2+(JumpTime*2)+(JumpTime-ps.playerEntity.jumpDuration2Step) {
			if ps.playerEntity.jumpPositionY > int(ps.playerEntity.SpaceComponent.Position.Y) {
				ps.playerEntity.SpaceComponent.Position.Y += JumpHeight
			}
		} else {
			ps.playerEntity.jumpDuration = 0
			ps.playerEntity.jumpCount = 0
		}
	}
	// GOAL
	if int(ps.playerEntity.SpaceComponent.Position.X) >= GoalPoint && ps.playerEntity.ifGoal == false {
		ps.playerEntity.ifGoal = true
		whenGoal(ps)
	}
}

// New :
func (ps *PlayerSystem) New(w *ecs.World) {
	ps.world = w
	// プレーヤーの作成
	player := Player{BasicEntity: ecs.NewBasic()}

	// 初期の配置
	PsPositionX := engo.WindowWidth() / 2
	PsPositionY := float32(PlayerPositionY)

	// SpaceComponent
	player.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: PsPositionX, Y: PsPositionY},
		Width:    30,
		Height:   30,
	}
	po := (int)(PsPositionX)
	fmt.Println("プレイヤー位置X:" + strconv.Itoa(po))
	po = (int)(PsPositionY)
	fmt.Println("プレイヤー位置Y:" + strconv.Itoa(po))

	// スプライトシートの作成
	player.Spritesheet = common.NewSpritesheetWithBorderFromFile(playerFile, 16, 16, 0, 0)
	// 画像の読み込み
	/*texture, err := common.LoadedSprite("gopher.png")
	if err != nil {
		fmt.Println("Unable to load texture: " + err.Error())
	}*/
	// RenderComponent
	player.RenderComponent = common.RenderComponent{
		Drawable: player.Spritesheet.Cell(SpriteSheetCell),
		Scale:    engo.Point{X: 2, Y: 2},
	}
	player.RenderComponent.SetZIndex(1001)

	// Entity
	ps.playerEntity = &player

	// カメラ位置
	ps.playerEntity.distance = 0
	// ダメージ初期化
	ps.playerEntity.damage = 0

	for _, system := range ps.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&player.BasicEntity, &player.RenderComponent, &player.SpaceComponent)
		}
	}
	common.CameraBounds = engo.AABB{
		Min: engo.Point{X: 0, Y: 0},
		Max: engo.Point{X: 40000, Y: 300},
	}
	fmt.Println("カメラ位置：" + strconv.Itoa(ps.playerEntity.distance))
}
func whenGoal(ps *PlayerSystem) {
	for _, system := range ps.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Remove(ps.playerEntity.BasicEntity)
		case *HUDTextSystem:
			AddText(ps.world, sys, TextGOAL)
			AddText(ps.world, sys, TextRETRY)
		case *ButtonSystem:
			//AddButton(ps.world, sys, BtnEND)
			AddButton(ps.world, sys, BtnRETRY)
		}
	}
}
func whenDied(ps *PlayerSystem) {
	for _, system := range ps.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Remove(ps.playerEntity.BasicEntity)
		case *HUDTextSystem:
			AddText(ps.world, sys, TextEND)
			AddText(ps.world, sys, TextRETRY)
		case *ButtonSystem:
			//AddButton(ps.world, sys, BtnEND)
			AddButton(ps.world, sys, BtnRETRY)
		}
	}
}
