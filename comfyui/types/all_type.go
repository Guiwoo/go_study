package types

type Prompt struct {
	Msg       string `json:"prompt"`
	Translate string `json:"translate"`
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
	Type string `json:"type"`
	//Data ComfySocketData `json:"data"`
}
type ComfySocketData struct {
	Node   string             `json:"node"`
	Images []ComfySocketImage `json:"images"`
}
type ComfySocketImage struct {
	FileName string `json:"filename"`
	Type     string `json:"type"`
}
