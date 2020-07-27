package generator

import (
	"fmt"
	"math/rand"
	"time"
)

func Generate() {
	components := []string{"database", "core", "middleware", "auth", "queue"}
	logLevels := []string{"DEBUG", "INFO", "WARNING", "ERROR", "CRITICAL"}

	// message for database components
	dbDebugMsgs := []string{"connected to datbase",
		"graceful temination of connection",
	}
	dbInfoMsgs := []string{"connection successful", "database connection lateny is 5ms"}
	dbWarningMsgs := []string{
		"database latency is 100ms",
		"ungraceful connection termination",
		"CPU utilization gt 95 percent",
		"burst balance limit reached",
	}
	dbError := []string{"failed to connection to db", "timout connecting to db"}
	dbCritical := []string{"max_connections limit reached", "database crashed"}

	// core messages
  coreDebugMsg := []string{"Pushed job to celery queue", "Fetching query for cache"}
  coreWarningMsg := []string{"too many cache misses", "Un authenticated access"}

  // middleware messages
  middleWareInfoMsgs := []string{"connected to auth success"}
  middleWareDebugMsg := []string{"un authorized access"}
  middleWareWarningMsg := []string{"to many requests"}
  middleWareErrorMsg := []string{"to many requests"}
  middleWareCriticalMsg := []string{"to rebooting webserver"}

  // auth messages
  authInfoMsg := []string{"connected to auth"}

  // queue messages
  queueMsg := []string{"queue is full"}

	for true {
		component := components[rand.Intn(len(components)-0)+0]
		logLevel := logLevels[rand.Intn(len(logLevels)-0)+0]
    messageType := fmt.Sprintf("%s-%s", component, logLevel)
    var msg string
    switch messageType {
    case "database-DEBUG":
       msg = dbDebugMsgs[rand.Intn(len(dbDebugMsgs)-0)+0]
     case "database-INFO":
       msg = dbInfoMsgs[rand.Intn(len(dbInfoMsgs)-0)+0]
       case "database-WARNING":
       msg = dbWarningMsgs[rand.Intn(len(dbWarningMsgs)-0)+0]
     case "database-ERROR":
       msg = dbError[rand.Intn(len(dbError)-0)+0]
    case "database-CRITICAL":
      msg = dbCritical[rand.Intn(len(dbInfoMsgs)-0)+0]
    case "core-DEBUG":
      msg = coreDebugMsg[rand.Intn(len(dbInfoMsgs)-0)+0]
    case "core-WARNING":
      msg = coreWarningMsg[rand.Intn(len(dbInfoMsgs)-0)+0]
    case "auth-INFO":
      msg = authInfoMsg[rand.Intn(len(authInfoMsg)-0)+0]
    case "queue-INFO":
      msg = queueMsg[rand.Intn(len(queueMsg)-0)+0]
    case "middlware-DEBUG":
      msg = middleWareDebugMsg[rand.Intn(len(middleWareDebugMsg)-0)+0]
    case "middlware-INFO":
      msg = middleWareInfoMsgs[rand.Intn(len(middleWareInfoMsgs)-0)+0]
      case "middlware-WARNING":
      msg = middleWareWarningMsg[rand.Intn(len(middleWareWarningMsg)-0)+0]
    case "middlware-ERROR":
      msg = middleWareErrorMsg[rand.Intn(len(middleWareErrorMsg)-0)+0]
   case "middlware-CRITICAL":
     msg = middleWareCriticalMsg[rand.Intn(len(middleWareCriticalMsg)-0)+0]
    default:
      msg = "connected to db"
  }
		date := time.Now().Format("01-02-2006 15:04:05.000000")
		fmt.Printf("%s %10s %10s %s\n", date, component, logLevel, msg)
	}
}
