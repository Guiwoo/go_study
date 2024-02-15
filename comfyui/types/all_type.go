package types

type Prompt struct {
	Positive   string `json:"positive"`
	Negative   string `json:"negative"`
	PTranslate string `json:"p_translate"`
	NTranslate string `json:"n_translate"`
}

type PapagoResp struct {
	Message PapagoMsg `json:"message"`
}

type PapagoMsg struct {
	Type    string       `json:"@type"`
	Service string       `json:"@service"`
	Version string       `json:"@version"`
	Result  PapagoResult `json:"result"`
}

type PapagoResult struct {
	SrcLangType    string `json:"srcLangType"`
	TarLangType    string `json:"tarLangType"`
	TranslatedText string `json:"translatedText"`
	EngineType     string `json:"engineType"`
}

type ComfySocketResp struct {
	Type string          `json:"type"`
	Data ComfySocketData `json:"data"`
}
type ComfySocketData struct {
	Node   string      `json:"node"`
	Output ComfyOutput `json:"output"`
}
type ComfyOutput struct {
	Images []ComfySocketImage `json:"images"`
}

type ComfySocketImage struct {
	FileName string `json:"filename"`
	Type     string `json:"type"`
}

type QueueRequest struct {
	Model     string `json:"model"`
	Positive  string `json:"positive"`
	Negative  string `json:"negative"`
	Seed      string `json:"seed"`
	Cfg       string `json:"cfg"`
	Steps     string `json:"steps"`
	Width     string `json:"width"`
	Height    string `json:"height"`
	BatchSize string `json:"batchSize"`
}
