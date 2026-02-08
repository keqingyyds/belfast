package answer

import (
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/connection"

	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/protobuf"
)

func GetMetaShipsPointsResponse(buffer *[]byte, client *connection.Client) (int, int, error) {
	var response protobuf.SC_34002
	return client.SendMessage(34002, &response)
}
