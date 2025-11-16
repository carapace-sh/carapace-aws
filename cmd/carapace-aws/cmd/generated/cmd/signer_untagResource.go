package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var signer_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from a signing profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(signer_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(signer_untagResourceCmd).Standalone()

		signer_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for the signing profile.")
		signer_untagResourceCmd.Flags().String("tag-keys", "", "A list of tag keys to be removed from the signing profile.")
		signer_untagResourceCmd.MarkFlagRequired("resource-arn")
		signer_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	signerCmd.AddCommand(signer_untagResourceCmd)
}
