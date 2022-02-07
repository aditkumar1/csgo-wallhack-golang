package main

import (
	"github.com/aditkumar1/goxymemory"
	"log"
	"os"
)

const(
	defaultDwEntityListOffset        = 81587068
	defaultDwGlowObjectManagerOffset = 87125656
	defaultMIGlowIndexOffset         = 66696
	defaultMIteamnumOffset           = 244
	enableVal                        = 1
	disableVal                       = 0
)

func main(){
	// panic recovery
	defer func(){ if r := recover(); r!= nil {
		log.Printf("Encountered error - %v , ensure CS GO application running before starting this application.\n",r)
		log.Println("Exiting application")
		os.Exit(1)
	}}()
	log.Printf("launching cheat ......")
	dm := goxymemmory.DataManager("csgo.exe")
	clientAddress ,_ := dm.GetModuleFromName("client.dll")
	glowManager, _ := dm.Read((uint)(clientAddress) + defaultDwGlowObjectManagerOffset, goxymemmory.UINT)
	for{
		for i := 1; i<32; i++{
			entity, _ := dm.Read((uint)(clientAddress) + (uint)(defaultDwEntityListOffset+ i * 0x10), goxymemmory.UINT)
			if entity.Value.(uint32) > 0{
				entityTeamId,_ := dm.Read((uint)(entity.Value.(uint32) + defaultMIteamnumOffset),goxymemmory.UINT)
				entityGlow,_ := dm.Read((uint)(entity.Value.(uint32) + defaultMIGlowIndexOffset),goxymemmory.UINT)
				red := (uint)(glowManager.Value.(uint32) + entityGlow.Value.(uint32) * 0x38 + 0x8)
				green := (uint)(glowManager.Value.(uint32) + entityGlow.Value.(uint32) * 0x38 + 0xC)
				blue := (uint)(glowManager.Value.(uint32) + entityGlow.Value.(uint32) * 0x38 + 0x10)
				alpha := (uint)(glowManager.Value.(uint32) + entityGlow.Value.(uint32) * 0x38 + 0x14)
				enable := (uint)(glowManager.Value.(uint32) + entityGlow.Value.(uint32) * 0x38 + 0x28)
				if entityTeamId.Value.(uint32) == 2 {
					dm.Write(red,
						goxymemmory.Data{Value:float32(enableVal), DataType:goxymemmory.FLOAT})
					dm.Write(green,
						goxymemmory.Data{Value:float32(disableVal), DataType:goxymemmory.FLOAT})
					dm.Write(blue,
						goxymemmory.Data{Value:float32(disableVal), DataType:goxymemmory.FLOAT})
					dm.Write(alpha,
						goxymemmory.Data{Value:float32(enableVal), DataType:goxymemmory.FLOAT})
					dm.Write(enable,
						goxymemmory.Data{Value:enableVal, DataType:goxymemmory.UINT})
				}
				if entityTeamId.Value.(uint32) == 3 {
					dm.Write(red,
						goxymemmory.Data{Value:float32(disableVal), DataType:goxymemmory.FLOAT})
					dm.Write(green,
						goxymemmory.Data{Value:float32(disableVal), DataType:goxymemmory.FLOAT})
					dm.Write(blue,
						goxymemmory.Data{Value:float32(enableVal), DataType:goxymemmory.FLOAT})
					dm.Write(alpha,
						goxymemmory.Data{Value:float32(enableVal), DataType:goxymemmory.FLOAT})
					dm.Write(enable,
						goxymemmory.Data{Value:enableVal, DataType:goxymemmory.UINT})
				}
			}
		}

	}
}

