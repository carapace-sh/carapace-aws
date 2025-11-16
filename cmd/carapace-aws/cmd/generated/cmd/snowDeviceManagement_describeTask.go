package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowDeviceManagement_describeTaskCmd = &cobra.Command{
	Use:   "describe-task",
	Short: "Checks the metadata for a given task on a device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowDeviceManagement_describeTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(snowDeviceManagement_describeTaskCmd).Standalone()

		snowDeviceManagement_describeTaskCmd.Flags().String("task-id", "", "The ID of the task to be described.")
		snowDeviceManagement_describeTaskCmd.MarkFlagRequired("task-id")
	})
	snowDeviceManagementCmd.AddCommand(snowDeviceManagement_describeTaskCmd)
}
