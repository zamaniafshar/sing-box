package option

import (
	"github.com/sagernet/sing/common/json/badjson"
)

type SSMAPIServiceOptions struct {
	ListenOptions
	InboundTLSOptionsContainer
	Servers *badjson.TypedMap[string, string] `json:"servers"`
}
