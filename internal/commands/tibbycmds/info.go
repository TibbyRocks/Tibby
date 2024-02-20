package tibbycmds

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/tibbyrocks/tibby/internal/utils"
)

var version string
var appStart time.Time

func GetInfo(i *discordgo.InteractionCreate, s *discordgo.Session) string {
	utils.LogCmd(i)
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	uptime := time.Since(appStart).Truncate(time.Second).String()
	return fmt.Sprintf(infoFormat, version, uptime, appStart.Format("02-01-2006 15:04:05 MST"), os.Getenv("WB_TRANSLATOR"), os.Getenv("WB_LANGUAGELOOKUP"), bToMb(m.Alloc), bToMb(m.TotalAlloc), bToMb(m.Sys), s.HeartbeatLatency().Milliseconds())
}

func RegisterAppStart() {
	appStart = time.Now()
}

var infoFormat string = `
	Application Version: %s
	Uptime: %s (since %s)
	
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
