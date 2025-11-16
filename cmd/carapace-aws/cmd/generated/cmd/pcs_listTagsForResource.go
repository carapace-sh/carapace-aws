package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcs_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of all tags on an PCS resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcs_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcs_listTagsForResourceCmd).Standalone()

		pcs_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource for which to list tags.")
		pcs_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	pcsCmd.AddCommand(pcs_listTagsForResourceCmd)
}
