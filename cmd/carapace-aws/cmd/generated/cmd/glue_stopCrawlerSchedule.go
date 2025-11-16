package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_stopCrawlerScheduleCmd = &cobra.Command{
	Use:   "stop-crawler-schedule",
	Short: "Sets the schedule state of the specified crawler to `NOT_SCHEDULED`, but does not stop the crawler if it is already running.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_stopCrawlerScheduleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_stopCrawlerScheduleCmd).Standalone()

		glue_stopCrawlerScheduleCmd.Flags().String("crawler-name", "", "Name of the crawler whose schedule state to set.")
		glue_stopCrawlerScheduleCmd.MarkFlagRequired("crawler-name")
	})
	glueCmd.AddCommand(glue_stopCrawlerScheduleCmd)
}
