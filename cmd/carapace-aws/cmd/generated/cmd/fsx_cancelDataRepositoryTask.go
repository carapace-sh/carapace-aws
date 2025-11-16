package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_cancelDataRepositoryTaskCmd = &cobra.Command{
	Use:   "cancel-data-repository-task",
	Short: "Cancels an existing Amazon FSx for Lustre data repository task if that task is in either the `PENDING` or `EXECUTING` state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_cancelDataRepositoryTaskCmd).Standalone()

	fsx_cancelDataRepositoryTaskCmd.Flags().String("task-id", "", "Specifies the data repository task to cancel.")
	fsx_cancelDataRepositoryTaskCmd.MarkFlagRequired("task-id")
	fsxCmd.AddCommand(fsx_cancelDataRepositoryTaskCmd)
}
