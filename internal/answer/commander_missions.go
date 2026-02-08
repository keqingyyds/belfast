package answer

import (
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/connection"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/protobuf"
)

func CommanderMissions(buffer *[]byte, client *connection.Client) (int, int, error) {
	var response protobuf.SC_20001
	return client.SendMessage(20001, &response)
}
