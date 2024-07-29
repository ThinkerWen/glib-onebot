package messages

type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type TextData struct {
	Text string `json:"text"`
}

type FaceData struct {
	Id string `json:"id"`
}

type ImageData struct {
	File    string `json:"file"`
	Type    string `json:"type"`
	Url     string `json:"url"`
	Cache   int    `json:"cache"`
	Proxy   int    `json:"proxy"`
	Timeout int    `json:"timeout"`
}

type RecordData struct {
	File    string `json:"file"`
	Magic   int    `json:"magic"`
	Url     string `json:"url"`
	Cache   int    `json:"cache"`
	Proxy   int    `json:"proxy"`
	Timeout int    `json:"timeout"`
}

type VideoData struct {
	File    string `json:"file"`
	Url     string `json:"url"`
	Cache   int    `json:"cache"`
	Proxy   int    `json:"proxy"`
	Timeout int    `json:"timeout"`
}

type AtData struct {
	QQ string `json:"qq"`
}

type RpsData struct {
}

type DiceData struct {
}

type ShakeData struct {
}

type PokeData struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	Name string `json:"name"`
}

type AnonymousData struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Flag string `json:"flag"`
}

type ShareData struct {
	Url     string `json:"url"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

type ContactData struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

type LocationData struct {
	Lat     string `json:"lat"`
	Lon     string `json:"lon"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type MusicData struct {
	Type    string `json:"type"`
	Id      string `json:"id"`
	Url     string `json:"url"`
	Audio   string `json:"audio"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

type ReplyData struct {
	Id string `json:"id"`
}

type ForwardData struct {
	Id string `json:"id"`
}

type NodeData struct {
	Id       string  `json:"id"`
	UserId   string  `json:"user_id"`
	Nickname string  `json:"nickname"`
	Content  Message `json:"content"`
}

type XmlData struct {
	Data string `json:"data"`
}

type JsonData struct {
	Data string `json:"data"`
}
