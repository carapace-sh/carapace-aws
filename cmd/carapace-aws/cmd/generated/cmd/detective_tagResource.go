package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detective_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Applies tag values to a behavior graph.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detective_tagResourceCmd).Standalone()

	detective_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the behavior graph to assign the tags to.")
	detective_tagResourceCmd.Flags().String("tags", "", "The tags to assign to the behavior graph.")
	detective_tagResourceCmd.MarkFlagRequired("resource-arn")
	detective_tagResourceCmd.MarkFlagRequired("tags")
	detectiveCmd.AddCommand(detective_tagResourceCmd)
}
