package answer

import (
	"errors"

	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/connection"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/orm"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/protobuf"
	"google.golang.org/protobuf/proto"
)

func JuustagramReadTip(buffer *[]byte, client *connection.Client) (int, int, error) {
	var payload protobuf.CS_11720
	if err := proto.Unmarshal(*buffer, &payload); err != nil {
		return 0, 11721, err
	}
	if client.Commander == nil {
		return 0, 11721, errors.New("missing commander")
	}
	if err := orm.MarkJuustagramChatGroupsRead(client.Commander.CommanderID, payload.GetChatGroupIdList()); err != nil {
		return 0, 11721, err
	}
	response := protobuf.SC_11721{
		Result: proto.Uint32(0),
	}
	return client.SendMessage(11721, &response)
}
