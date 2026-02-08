package answer

import (
	"encoding/json"

	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/connection"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/orm"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/protobuf"
	"google.golang.org/protobuf/proto"
)

type permanentActivity struct {
	ID uint32 `json:"id"`
}

func PermanentActivites(buffer *[]byte, client *connection.Client) (int, int, error) {
	entries, err := orm.ListConfigEntries(orm.GormDB, "ShareCfg/activity_task_permanent.json")
	if err != nil {
		return 0, 11210, err
	}
	state, err := orm.GetOrCreateActivityPermanentState(orm.GormDB, client.Commander.CommanderID)
	if err != nil {
		return 0, 11210, err
	}
	response := protobuf.SC_11210{
		PermanentActivity: make([]uint32, 0, len(entries)),
		PermanentNow:      proto.Uint32(state.CurrentActivityID),
	}
	for _, entry := range entries {
		var activity permanentActivity
		if err := json.Unmarshal(entry.Data, &activity); err != nil {
			return 0, 11210, err
		}
		response.PermanentActivity = append(response.PermanentActivity, activity.ID)
	}
	return client.SendMessage(11210, &response)
}
