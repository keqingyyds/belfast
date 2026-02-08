package answer

import (
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/connection"

	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/protobuf"
)

func MetaCharacterTacticsInfoRequestCommandResponse(buffer *[]byte, client *connection.Client) (int, int, error) {
	var response protobuf.SC_63318
	return client.SendMessage(63318, &response)
}
