package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowDeviceManagement_createTaskCmd = &cobra.Command{
	Use:   "create-task",
	Short: "Instructs one or more devices to start a task, such as unlocking or rebooting.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowDeviceManagement_createTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(snowDeviceManagement_createTaskCmd).Standalone()

		snowDeviceManagement_createTaskCmd.Flags().String("client-token", "", "A token ensuring that the action is called only once with the specified details.")
		snowDeviceManagement_createTaskCmd.Flags().String("command", "", "The task to be performed.")
		snowDeviceManagement_createTaskCmd.Flags().String("description", "", "A description of the task and its targets.")
		snowDeviceManagement_createTaskCmd.Flags().String("tags", "", "Optional metadata that you assign to a resource.")
		snowDeviceManagement_createTaskCmd.Flags().String("targets", "", "A list of managed device IDs.")
		snowDeviceManagement_createTaskCmd.MarkFlagRequired("command")
		snowDeviceManagement_createTaskCmd.MarkFlagRequired("targets")
	})
	snowDeviceManagementCmd.AddCommand(snowDeviceManagement_createTaskCmd)
}
