package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of the tags assigned to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_listTagsForResourceCmd).Standalone()

		ssm_listTagsForResourceCmd.Flags().String("resource-id", "", "The resource ID for which you want to see a list of tags.")
		ssm_listTagsForResourceCmd.Flags().String("resource-type", "", "Returns a list of tags for a specific resource type.")
		ssm_listTagsForResourceCmd.MarkFlagRequired("resource-id")
		ssm_listTagsForResourceCmd.MarkFlagRequired("resource-type")
	})
	ssmCmd.AddCommand(ssm_listTagsForResourceCmd)
}
