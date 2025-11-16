package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var signer_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds one or more tags to a signing profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(signer_tagResourceCmd).Standalone()

	signer_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for the signing profile.")
	signer_tagResourceCmd.Flags().String("tags", "", "One or more tags to be associated with the signing profile.")
	signer_tagResourceCmd.MarkFlagRequired("resource-arn")
	signer_tagResourceCmd.MarkFlagRequired("tags")
	signerCmd.AddCommand(signer_tagResourceCmd)
}
