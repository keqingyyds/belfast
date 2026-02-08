package answer

import (
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/connection"

	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/protobuf"
	"google.golang.org/protobuf/proto"
)

func Meowfficers(buffer *[]byte, client *connection.Client) (int, int, error) {
	var response protobuf.SC_25001
	response.UsageCount = proto.Uint32(0)
	return client.SendMessage(25001, &response)
}
