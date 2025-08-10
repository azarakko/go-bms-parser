package bmson

// Bmson はbmsonファイルのルートオブジェクトです。
type Bmson struct {
	Version       string         `json:"version"`                 // bmsonフォーマットのバージョン
	Info          Info           `json:"info"`                    // 譜面のメタデータ
	Lines         []BarLine      `json:"lines"`                   // 小節線の位置
	BpmEvents     []BpmEvent     `json:"bpm_events"`              // BPM変更イベント
	StopEvents    []StopEvent    `json:"stop_events"`             // 停止イベント
	ScrollEvents  []ScrollEvent  `json:"scroll_events,omitempty"` // スクロール速度変更イベント
	SoundChannels []SoundChannel `json:"sound_channels"`          // ノートデータ
	Bga           *BGA           `json:"bga,omitempty"`           // BGAデータ
	MineChannels  []MineChannel  `json:"mine_channels,omitempty"` // 地雷ノートデータ
	KeyChannels   []KeyChannel   `json:"key_channels,omitempty"`  // キーサウンドノートデータ
}

// Info は譜面のメタデータを格納します。
type Info struct {
	Title         string   `json:"title"`
	Subtitle      string   `json:"subtitle"`
	Artist        string   `json:"artist"`
	SubArtists    []string `json:"subartists"`
	Genre         string   `json:"genre"`
	ModeHint      string   `json:"mode_hint"`      // レイアウトのヒント (例: "beat-7k")
	ChartName     string   `json:"chart_name"`     // チャート名 (例: "HYPER")
	Level         int      `json:"level"`          // 難易度レベル
	InitBPM       float64  `json:"init_bpm"`       // 初期BPM
	JudgeRank     float64  `json:"judge_rank"`     // 判定の厳しさ
	Total         float64  `json:"total"`          // ゲージ増加量
	BackImage     *string  `json:"back_image"`     // 背景画像ファイル名
	EyecatchImage *string  `json:"eyecatch_image"` // アイキャッチ画像ファイル名
	BannerImage   *string  `json:"banner_image"`   // バナー画像ファイル名
	PreviewMusic  *string  `json:"preview_music"`  // プレビュー音楽ファイル名
	Resolution    int      `json:"resolution"`     // 1拍あたりのパルス数
}

// BarLine は小節線の位置を表します。
type BarLine struct {
	Y int `json:"y"` // パルス位置
}

// BpmEvent はBPM変更イベントを表します。
type BpmEvent struct {
	Y   int     `json:"y"`   // パルス位置
	BPM float64 `json:"bpm"` // 適用される新しいBPM値
}

// StopEvent は停止イベントを表します。
type StopEvent struct {
	Y        int `json:"y"`        // パルス位置
	Duration int `json:"duration"` // 停止するパルス数
}

// ScrollEvent はスクロール速度変更イベントを表します。
type ScrollEvent struct {
	Y    int     `json:"y"`    // パルス位置
	Rate float64 `json:"rate"` // 適用される新しいスクロール倍率
}

// SoundChannel は特定の音声ファイルに対応するノート群を格納します。
type SoundChannel struct {
	Name  string `json:"name"`  // 音声ファイル名
	Notes []Note `json:"notes"` // このチャンネルに属するノートのリスト
}

// Note は単一のプレイ可能なノートを表します。
type Note struct {
	X int  `json:"x"` // レーン番号
	Y int  `json:"y"` // パルス位置
	L int  `json:"l"` // ノートの長さ（パルス単位）。0は通常ノート。
	C bool `json:"c"` // ロングノートの終端を示す継続フラグ
}

// BGA は全てのBGA関連イベントと定義を格納します。
type BGA struct {
	BGAHeader   []BGAHeader `json:"bga_header"`   // 画像IDとファイル名のマッピング
	BGAEvents   []BGAEvent  `json:"bga_events"`   // 通常のBGA表示イベント
	LayerEvents []BGAEvent  `json:"layer_events"` // レイヤーBGA表示イベント
	PoorEvents  []BGAEvent  `json:"poor_events"`  // POOR判定時のBGA表示イベント
}

// BGAHeader はBGAのIDとファイル名をマッピングします。
type BGAHeader struct {
	ID   int    `json:"id"`   // 画像の識別子
	Name string `json:"name"` // 画像ファイル名
}

// BGAEvent はBGA表示イベントを表します。
type BGAEvent struct {
	Y  int `json:"y"`  // パルス位置
	ID int `json:"id"` // 表示する画像のID
}

// MineChannel は地雷ノートを格納します。
type MineChannel struct {
	Notes []MineNote `json:"notes"`
}

// MineNote はダメージを発生させる地雷ノートを表します。
type MineNote struct {
	X      int `json:"x"`      // レーン番号
	Y      int `json:"y"`      // パルス位置
	Damage int `json:"damage"` // 地雷によるダメージ量
}

// KeyChannel はキーサウンドノートを格納します。
type KeyChannel struct {
	Notes []KeyNote `json:"notes"`
}

// KeyNote はキーサウンドノートを表します。
type KeyNote struct {
	X   int    `json:"x"`   // レーン番号
	Y   int    `json:"y"`   // パルス位置
	Key string `json:"key"` // キーサウンドのID
}
