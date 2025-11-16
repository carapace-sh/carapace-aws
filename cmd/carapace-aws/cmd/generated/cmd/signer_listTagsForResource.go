package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var signer_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of the tags associated with a signing profile resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(signer_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(signer_listTagsForResourceCmd).Standalone()

		signer_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for the signing profile.")
		signer_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	signerCmd.AddCommand(signer_listTagsForResourceCmd)
}
