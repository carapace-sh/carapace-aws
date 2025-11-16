package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_cancelExportTaskCmd = &cobra.Command{
	Use:   "cancel-export-task",
	Short: "Cancels an export task in progress that is exporting a snapshot or cluster to Amazon S3.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_cancelExportTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_cancelExportTaskCmd).Standalone()

		rds_cancelExportTaskCmd.Flags().String("export-task-identifier", "", "The identifier of the snapshot or cluster export task to cancel.")
		rds_cancelExportTaskCmd.MarkFlagRequired("export-task-identifier")
	})
	rdsCmd.AddCommand(rds_cancelExportTaskCmd)
}
