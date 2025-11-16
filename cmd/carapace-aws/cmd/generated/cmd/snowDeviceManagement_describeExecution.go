package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowDeviceManagement_describeExecutionCmd = &cobra.Command{
	Use:   "describe-execution",
	Short: "Checks the status of a remote task running on one or more target devices.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowDeviceManagement_describeExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(snowDeviceManagement_describeExecutionCmd).Standalone()

		snowDeviceManagement_describeExecutionCmd.Flags().String("managed-device-id", "", "The ID of the managed device.")
		snowDeviceManagement_describeExecutionCmd.Flags().String("task-id", "", "The ID of the task that the action is describing.")
		snowDeviceManagement_describeExecutionCmd.MarkFlagRequired("managed-device-id")
		snowDeviceManagement_describeExecutionCmd.MarkFlagRequired("task-id")
	})
	snowDeviceManagementCmd.AddCommand(snowDeviceManagement_describeExecutionCmd)
}
