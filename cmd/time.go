package cmd

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/zhengzhou1992/cmdtools/internal/timer"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "time formater",
	Long:  "time formater",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "get now time",
	Long:  "get now time",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("output: %s, %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

var startTimeStr string
var duration string

var calTimeCommand = &cobra.Command{
	Use:   "calc",
	Short: "calculate time",
	Long:  "calculate time",
	Run: func(cmd *cobra.Command, args []string) {
		var startTime time.Time
		var layout = "2006-01-02 15:04:05"
		var err error
		if startTimeStr == "" {
			startTime = timer.GetNowTime()
		} else {
			space := strings.Count(startTimeStr, " ")
			if space == 0 {
				layout = "2020-08-08"
			} else {
				layout = "2020-08-08 17:00"
			}
			startTime, err = time.Parse(layout, startTimeStr)
			if err != nil {
				t, _ := strconv.Atoi(startTimeStr)
				startTime = time.Unix(int64(t), 0)
			}
		}
		t, err := timer.GetCalculateTime(startTime, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err: %v", err)
		}
		log.Printf("Output: %s, %d", t.Format(layout), t.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calTimeCommand)

	calTimeCommand.Flags().StringVarP(&startTimeStr, "start time", "s", "", "the start time or timestamp")
	calTimeCommand.Flags().StringVarP(&duration, "duration", "d", "", `duration, supoort unit: "ns", "us" (or "µs"), "ms", "s", "m", "h"`)
}
