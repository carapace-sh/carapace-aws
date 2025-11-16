package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns one or more tags (key-value pairs) to the specified Amazon Quick Sight resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_tagResourceCmd).Standalone()

		quicksight_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to tag.")
		quicksight_tagResourceCmd.Flags().String("tags", "", "Contains a map of the key-value pairs for the resource tag or tags assigned to the resource.")
		quicksight_tagResourceCmd.MarkFlagRequired("resource-arn")
		quicksight_tagResourceCmd.MarkFlagRequired("tags")
	})
	quicksightCmd.AddCommand(quicksight_tagResourceCmd)
}
