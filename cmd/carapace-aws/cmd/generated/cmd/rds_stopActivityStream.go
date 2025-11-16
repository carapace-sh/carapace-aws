package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_stopActivityStreamCmd = &cobra.Command{
	Use:   "stop-activity-stream",
	Short: "Stops a database activity stream that was started using the Amazon Web Services console, the `start-activity-stream` CLI command, or the `StartActivityStream` operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_stopActivityStreamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_stopActivityStreamCmd).Standalone()

		rds_stopActivityStreamCmd.Flags().String("apply-immediately", "", "Specifies whether or not the database activity stream is to stop as soon as possible, regardless of the maintenance window for the database.")
		rds_stopActivityStreamCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the DB cluster for the database activity stream.")
		rds_stopActivityStreamCmd.MarkFlagRequired("resource-arn")
	})
	rdsCmd.AddCommand(rds_stopActivityStreamCmd)
}
