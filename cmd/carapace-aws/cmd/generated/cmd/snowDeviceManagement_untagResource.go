package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowDeviceManagement_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag from a device or task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowDeviceManagement_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(snowDeviceManagement_untagResourceCmd).Standalone()

		snowDeviceManagement_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the device or task.")
		snowDeviceManagement_untagResourceCmd.Flags().String("tag-keys", "", "Optional metadata that you assign to a resource.")
		snowDeviceManagement_untagResourceCmd.MarkFlagRequired("resource-arn")
		snowDeviceManagement_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	snowDeviceManagementCmd.AddCommand(snowDeviceManagement_untagResourceCmd)
}
