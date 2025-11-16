package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var invoicing_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds a tag to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(invoicing_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(invoicing_tagResourceCmd).Standalone()

		invoicing_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the tags.")
		invoicing_tagResourceCmd.Flags().String("resource-tags", "", "Adds a tag to a resource.")
		invoicing_tagResourceCmd.MarkFlagRequired("resource-arn")
		invoicing_tagResourceCmd.MarkFlagRequired("resource-tags")
	})
	invoicingCmd.AddCommand(invoicing_tagResourceCmd)
}
