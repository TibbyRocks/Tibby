package tibbycmds

import (
	"fmt"
	"os"
	"runtime"

	"github.com/bwmarrin/discordgo"
)

var version string

func GetInfo(i *discordgo.InteractionCreate, s *discordgo.Session) string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return fmt.Sprintf(infoFormat, version, os.Getenv("WB_TRANSLATOR"), os.Getenv("WB_LANGUAGELOOKUP"), bToMb(m.Alloc), bToMb(m.TotalAlloc), bToMb(m.Sys), s.HeartbeatLatency().Milliseconds())
}

var infoFormat string = `
	Application Version: %s
	
	**Translations**
	Translator: %s
	Language Resolver: %s

	**Memory usage**
	Allocated Memory: %dMB
	Total Allocated: %dMB
	Reserved: %dMB

	**Discord**
	Heartbeat latency: %dms
`

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
