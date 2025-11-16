package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowDeviceManagement_cancelTaskCmd = &cobra.Command{
	Use:   "cancel-task",
	Short: "Sends a cancel request for a specified task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowDeviceManagement_cancelTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(snowDeviceManagement_cancelTaskCmd).Standalone()

		snowDeviceManagement_cancelTaskCmd.Flags().String("task-id", "", "The ID of the task that you are attempting to cancel.")
		snowDeviceManagement_cancelTaskCmd.MarkFlagRequired("task-id")
	})
	snowDeviceManagementCmd.AddCommand(snowDeviceManagement_cancelTaskCmd)
}
