package answer

import (
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/connection"

	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/protobuf"
	"google.golang.org/protobuf/proto"
)

func TechnologyRefreshList(buffer *[]byte, client *connection.Client) (int, int, error) {
	response := protobuf.SC_63000{
		RefreshFlag: proto.Uint32(0),
		Catchup: &protobuf.TECHNOLOGYCATCHUP{
			Version: proto.Uint32(0),
			Target:  proto.Uint32(0),
		},
	}
	return client.SendMessage(63000, &response)
}
