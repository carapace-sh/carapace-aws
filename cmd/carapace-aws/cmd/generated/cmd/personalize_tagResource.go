package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Add a list of tags to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_tagResourceCmd).Standalone()

	personalize_tagResourceCmd.Flags().String("resource-arn", "", "The resource's Amazon Resource Name (ARN).")
	personalize_tagResourceCmd.Flags().String("tags", "", "Tags to apply to the resource.")
	personalize_tagResourceCmd.MarkFlagRequired("resource-arn")
	personalize_tagResourceCmd.MarkFlagRequired("tags")
	personalizeCmd.AddCommand(personalize_tagResourceCmd)
}
