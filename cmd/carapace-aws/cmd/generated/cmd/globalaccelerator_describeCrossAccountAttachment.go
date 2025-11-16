package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_describeCrossAccountAttachmentCmd = &cobra.Command{
	Use:   "describe-cross-account-attachment",
	Short: "Gets configuration information about a cross-account attachment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_describeCrossAccountAttachmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_describeCrossAccountAttachmentCmd).Standalone()

		globalaccelerator_describeCrossAccountAttachmentCmd.Flags().String("attachment-arn", "", "The Amazon Resource Name (ARN) for the cross-account attachment to describe.")
		globalaccelerator_describeCrossAccountAttachmentCmd.MarkFlagRequired("attachment-arn")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_describeCrossAccountAttachmentCmd)
}
