package answer

import (
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/connection"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/protobuf"
	"google.golang.org/protobuf/proto"
)

func SellItem(buffer *[]byte, client *connection.Client) (int, int, error) {
	response := protobuf.SC_15009{
		Result: proto.Uint32(0),
	}
	return client.SendMessage(15009, &response)
}
