package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detective_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns the tag values that are assigned to a behavior graph.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detective_listTagsForResourceCmd).Standalone()

	detective_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the behavior graph for which to retrieve the tag values.")
	detective_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	detectiveCmd.AddCommand(detective_listTagsForResourceCmd)
}
