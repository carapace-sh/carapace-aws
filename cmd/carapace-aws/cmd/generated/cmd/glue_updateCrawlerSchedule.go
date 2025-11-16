package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_updateCrawlerScheduleCmd = &cobra.Command{
	Use:   "update-crawler-schedule",
	Short: "Updates the schedule of a crawler using a `cron` expression.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_updateCrawlerScheduleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_updateCrawlerScheduleCmd).Standalone()

		glue_updateCrawlerScheduleCmd.Flags().String("crawler-name", "", "The name of the crawler whose schedule to update.")
		glue_updateCrawlerScheduleCmd.Flags().String("schedule", "", "The updated `cron` expression used to specify the schedule (see [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html). For example, to run something every day at 12:15 UTC, you would specify: `cron(15 12 * * ?")
		glue_updateCrawlerScheduleCmd.MarkFlagRequired("crawler-name")
	})
	glueCmd.AddCommand(glue_updateCrawlerScheduleCmd)
}
