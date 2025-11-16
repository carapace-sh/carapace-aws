package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowDeviceManagement_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds or replaces tags on a device or task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowDeviceManagement_tagResourceCmd).Standalone()

	snowDeviceManagement_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the device or task.")
	snowDeviceManagement_tagResourceCmd.Flags().String("tags", "", "Optional metadata that you assign to a resource.")
	snowDeviceManagement_tagResourceCmd.MarkFlagRequired("resource-arn")
	snowDeviceManagement_tagResourceCmd.MarkFlagRequired("tags")
	snowDeviceManagementCmd.AddCommand(snowDeviceManagement_tagResourceCmd)
}
