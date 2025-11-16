package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_cancelExportTaskCmd = &cobra.Command{
	Use:   "cancel-export-task",
	Short: "Cancels an active export task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_cancelExportTaskCmd).Standalone()

	ec2_cancelExportTaskCmd.Flags().String("export-task-id", "", "The ID of the export task.")
	ec2_cancelExportTaskCmd.MarkFlagRequired("export-task-id")
	ec2Cmd.AddCommand(ec2_cancelExportTaskCmd)
}
