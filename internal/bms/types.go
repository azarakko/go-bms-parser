package bms

// PlayerMode はプレイヤーモード（例: シングル、ダブル）を定義します。
type PlayerMode int

const (
	PlayerSingle PlayerMode = 1 // シングルプレイ
	PlayerCouple PlayerMode = 2 // カップルプレイ
	PlayerDouble PlayerMode = 3 // ダブルプレイ
	PlayerBattle PlayerMode = 4 // バトルプレイ
)

// Rank は判定の厳しさを定義します。
type Rank int

const (
	RankVeryHard Rank = 0 // VERY HARD
	RankHard     Rank = 1 // HARD
	RankNormal   Rank = 2 // NORMAL
	RankEasy     Rank = 3 // EASY
	RankVeryEasy Rank = 4 // VERY EASY
)

// Header はBMSヘッダフィールドのメタデータを格納します。
type Header struct {
	Player       PlayerMode      `json:"player,omitempty"`
	Genre        string          `json:"genre,omitempty"`
	Title        string          `json:"title"`
	Subtitle     string          `json:"subtitle,omitempty"`
	Artist       string          `json`:"artist"`
	SubArtist    string          `json:"sub_artist,omitempty"`
	BPM          float64         `json:"bpm"`
	PlayLevel    int             `json:"play_level,omitempty"`
	Rank         Rank            `json:"rank,omitempty"`
	VolWav       int             `json:"vol_wav,omitempty"`
	Total        float64         `json:"total,omitempty"`
	StageFile    string          `json:"stage_file,omitempty"`
	Preview      string          `json:"preview,omitempty"`
	Lntype       int             `json:"lntype,omitempty"`
	Lnobj        string          `json:"lnobj,omitempty"`
	WavFiles     map[string]string `json:"wav_files"`
	BmpFiles     map[string]string `json:"bmp_files"`
	BgaFiles     map[string]string `json:"bga_files,omitempty"`
	StopSequence map[string]string `json:"stop_sequence,omitempty"` // ID -> 拍数
	BpmChanges   map[string]float64 `json:"bpm_changes,omitempty"`  // ID -> BPM
}

// Object はメインデータフィールドの単一のノートまたはイベントを表します。
// これは、特定のイベントタイプに解決される前の生の表現です。
type Object struct {
	Measure  int
	Channel  string
	Position float64 // 小節内の位置 (0.0から1.0)
	Value    string  // 2桁の36進数値
}

// TimeSignature は小節の拍子変更を表します (#xxx02)。
type TimeSignature struct {
	Measure int
	Value   float64 // 例: 4/4拍子の場合は1.0, 3/4拍子の場合は0.75
}

// BMS はパースされたBMSファイルのトップレベル構造です。
type BMS struct {
	Header         Header
	Objects        []Object
	TimeSignatures []TimeSignature
}

// NewHeader は慣例的なデフォルト値でヘッダを作成します。
func NewHeader() *Header {
	return &Header{
		Player:    PlayerSingle,
		BPM:       130.0,
		PlayLevel: 3,
		Rank:      RankNormal,
		VolWav:    100,
		WavFiles:  make(map[string]string),
		BmpFiles:  make(map[string]string),
		BgaFiles:  make(map[string]string),
		StopSequence: make(map[string]string),
		BpmChanges:   make(map[string]float64),
	}
}
