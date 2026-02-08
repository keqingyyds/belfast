package answer

import (
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/connection"

	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/protobuf"
)

func CommanderGuildTechnologies(buffer *[]byte, client *connection.Client) (int, int, error) {
	var response protobuf.SC_62101
	return client.SendMessage(62101, &response)
}
