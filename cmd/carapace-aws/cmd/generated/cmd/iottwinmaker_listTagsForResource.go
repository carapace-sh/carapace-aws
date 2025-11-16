package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all tags associated with a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iottwinmaker_listTagsForResourceCmd).Standalone()

		iottwinmaker_listTagsForResourceCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
		iottwinmaker_listTagsForResourceCmd.Flags().String("next-token", "", "The string that specifies the next page of results.")
		iottwinmaker_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
		iottwinmaker_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	iottwinmakerCmd.AddCommand(iottwinmaker_listTagsForResourceCmd)
}
