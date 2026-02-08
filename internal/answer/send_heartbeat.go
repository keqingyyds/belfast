package answer

import (
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/connection"

	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/protobuf"
	"google.golang.org/protobuf/proto"
)

func SendHeartbeat(buffer *[]byte, client *connection.Client) (int, int, error) {
	var response protobuf.SC_10101
	response.State = proto.Uint32(0)
	return client.SendMessage(10101, &response)
}
