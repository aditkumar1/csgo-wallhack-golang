package offset

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
)

const(
	defaultDwEntityListOffset        = 81583740
	defaultDwGlowObjectManagerOffset = 87122576
	defaultMIGlowIndexOffset         = 66696
	defaultMIteamnumOffset           = 244
	OffsetsURL = "https://raw.githubusercontent.com/frk1/hazedumper/master/csgo.json"
)

var jsonCodec  = jsoniter.ConfigFastest

type Offset struct {
	DwEntityListOffset uint
	DwGlowObjectManagerOffset   uint
	MIGlowIndexOffset uint
	MIteamNumOffset       uint
}

func GetNewOffset() (*Offset, error) {
	resp, err := http.Get(OffsetsURL)

	if err != nil {
		return &Offset{
			DwEntityListOffset:        defaultDwEntityListOffset,
			DwGlowObjectManagerOffset: defaultDwGlowObjectManagerOffset,
			MIGlowIndexOffset:         defaultMIteamnumOffset,
			MIteamNumOffset:           defaultMIteamnumOffset,
		}, fmt.Errorf("fail to get offset. Error - %v. Using default offsets. Cheat may not work",err)
	}

	defer resp.Body.Close()

	jsonBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return &Offset{
			DwEntityListOffset:        defaultDwEntityListOffset,
			DwGlowObjectManagerOffset: defaultDwGlowObjectManagerOffset,
			MIGlowIndexOffset:         defaultMIteamnumOffset,
			MIteamNumOffset:           defaultMIteamnumOffset,
		}, fmt.Errorf("unable to parse returned offset json. Error - %v. Using default offsets. Cheat may not work",err)
	}

	dwEntityList := jsoniter.Get(jsonBytes, "signatures","dwEntityList").ToUint()
	dwGlowObjectManager := jsoniter.Get(jsonBytes, "signatures","dwGlowObjectManager").ToUint()
	miGlowIndex := jsoniter.Get(jsonBytes, "netvars","m_iGlowIndex").ToUint()
	miTeamNum := jsoniter.Get(jsonBytes, "netvars","m_iTeamNum").ToUint()

	return &Offset{
		DwEntityListOffset:        dwEntityList,
		DwGlowObjectManagerOffset: dwGlowObjectManager,
		MIGlowIndexOffset:         miGlowIndex,
		MIteamNumOffset:           miTeamNum,
	}, nil
}
