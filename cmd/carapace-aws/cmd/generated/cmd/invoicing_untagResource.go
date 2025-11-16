package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var invoicing_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(invoicing_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(invoicing_untagResourceCmd).Standalone()

		invoicing_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) to untag.")
		invoicing_untagResourceCmd.Flags().String("resource-tag-keys", "", "Keys for the tags to be removed.")
		invoicing_untagResourceCmd.MarkFlagRequired("resource-arn")
		invoicing_untagResourceCmd.MarkFlagRequired("resource-tag-keys")
	})
	invoicingCmd.AddCommand(invoicing_untagResourceCmd)
}
