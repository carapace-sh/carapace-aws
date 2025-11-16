package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_createCrossAccountAttachmentCmd = &cobra.Command{
	Use:   "create-cross-account-attachment",
	Short: "Create a cross-account attachment in Global Accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_createCrossAccountAttachmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_createCrossAccountAttachmentCmd).Standalone()

		globalaccelerator_createCrossAccountAttachmentCmd.Flags().String("idempotency-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency—that is, the uniqueness—of the request.")
		globalaccelerator_createCrossAccountAttachmentCmd.Flags().String("name", "", "The name of the cross-account attachment.")
		globalaccelerator_createCrossAccountAttachmentCmd.Flags().String("principals", "", "The principals to include in the cross-account attachment.")
		globalaccelerator_createCrossAccountAttachmentCmd.Flags().String("resources", "", "The Amazon Resource Names (ARNs) for the resources to include in the cross-account attachment.")
		globalaccelerator_createCrossAccountAttachmentCmd.Flags().String("tags", "", "Add tags for a cross-account attachment.")
		globalaccelerator_createCrossAccountAttachmentCmd.MarkFlagRequired("idempotency-token")
		globalaccelerator_createCrossAccountAttachmentCmd.MarkFlagRequired("name")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_createCrossAccountAttachmentCmd)
}
