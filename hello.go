package main
  
import (
        log "github.com/sirupsen/logrus"
        "os"
        "time"
)

func init() {

        // Output to stdout instead of the default stderr
        // Can be any io.Writer, see below for File example
        log.SetOutput(os.Stdout)

        // Only log the warning severity or above.
        log.SetLevel(log.DebugLevel)
        log.SetFormatter(&log.JSONFormatter{
                //TimestampFormat: "RFC3339",
                DisableTimestamp: true,
                PrettyPrint: true,
                })
        // This can be removed if CPU overhead is too high
        log.SetReportCaller(true)

}

func main() {
        log.WithFields(log.Fields{
                "animal": "walrus",
                "size":   10,
                "timestamp": float64(time.Now().UnixNano())/1e9,
        }).Info("A group of walrus emerges from the ocean")

        log.WithFields(log.Fields{
                "omg":    true,
                "number": 122,
                "timestamp": float64(time.Now().UnixNano())/1e9,
        }).Warn("The group's number increased tremendously!")

        // A common pattern is to re-use fields between logging statements by re-using
        // the logrus.Entry returned from WithFields()
        contextLogger := log.WithFields(log.Fields{
                "common": "this is a common field",
                "other":  "I also should be logged always",
                "timestamp": float64(time.Now().UnixNano())/1e9,
        })

        contextLogger.Info("I'll be logged with common and other field")
        contextLogger.Info("Me too")
        contextLogger.WithFields(log.Fields{"added": []string{"new field","extra"}, "timestamp": float64(time.Now().UnixNano())/1e9}).Info("Now Me too")
}
