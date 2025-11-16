package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_deleteCrossAccountAttachmentCmd = &cobra.Command{
	Use:   "delete-cross-account-attachment",
	Short: "Delete a cross-account attachment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_deleteCrossAccountAttachmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_deleteCrossAccountAttachmentCmd).Standalone()

		globalaccelerator_deleteCrossAccountAttachmentCmd.Flags().String("attachment-arn", "", "The Amazon Resource Name (ARN) for the cross-account attachment to delete.")
		globalaccelerator_deleteCrossAccountAttachmentCmd.MarkFlagRequired("attachment-arn")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_deleteCrossAccountAttachmentCmd)
}
