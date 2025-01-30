package segments

import (
	"io"
	"net/http"
)

type MempoolBlock struct {
	base

	Height string
}

func (n *MempoolBlock) Enabled() bool {
	resp, err := http.Get("https://mempool.space/api/blocks/tip/height")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	n.Height = string(body)

	return true
}

func (n *MempoolBlock) Template() string {
	return " {{ .Height }} "
}
