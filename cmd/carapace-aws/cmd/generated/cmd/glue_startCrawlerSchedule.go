package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_startCrawlerScheduleCmd = &cobra.Command{
	Use:   "start-crawler-schedule",
	Short: "Changes the schedule state of the specified crawler to `SCHEDULED`, unless the crawler is already running or the schedule state is already `SCHEDULED`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_startCrawlerScheduleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_startCrawlerScheduleCmd).Standalone()

		glue_startCrawlerScheduleCmd.Flags().String("crawler-name", "", "Name of the crawler to schedule.")
		glue_startCrawlerScheduleCmd.MarkFlagRequired("crawler-name")
	})
	glueCmd.AddCommand(glue_startCrawlerScheduleCmd)
}
