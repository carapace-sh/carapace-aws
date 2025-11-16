package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowDeviceManagement_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of tags for a managed device or task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowDeviceManagement_listTagsForResourceCmd).Standalone()

	snowDeviceManagement_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the device or task.")
	snowDeviceManagement_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	snowDeviceManagementCmd.AddCommand(snowDeviceManagement_listTagsForResourceCmd)
}
