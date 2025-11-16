package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_updateCrossAccountAttachmentCmd = &cobra.Command{
	Use:   "update-cross-account-attachment",
	Short: "Update a cross-account attachment to add or remove principals or resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_updateCrossAccountAttachmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_updateCrossAccountAttachmentCmd).Standalone()

		globalaccelerator_updateCrossAccountAttachmentCmd.Flags().String("add-principals", "", "The principals to add to the cross-account attachment.")
		globalaccelerator_updateCrossAccountAttachmentCmd.Flags().String("add-resources", "", "The resources to add to the cross-account attachment.")
		globalaccelerator_updateCrossAccountAttachmentCmd.Flags().String("attachment-arn", "", "The Amazon Resource Name (ARN) of the cross-account attachment to update.")
		globalaccelerator_updateCrossAccountAttachmentCmd.Flags().String("name", "", "The name of the cross-account attachment.")
		globalaccelerator_updateCrossAccountAttachmentCmd.Flags().String("remove-principals", "", "The principals to remove from the cross-account attachment.")
		globalaccelerator_updateCrossAccountAttachmentCmd.Flags().String("remove-resources", "", "The resources to remove from the cross-account attachment.")
		globalaccelerator_updateCrossAccountAttachmentCmd.MarkFlagRequired("attachment-arn")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_updateCrossAccountAttachmentCmd)
}
