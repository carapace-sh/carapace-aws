package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_createFilterCmd = &cobra.Command{
	Use:   "create-filter",
	Short: "Creates a recommendation filter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_createFilterCmd).Standalone()

	personalize_createFilterCmd.Flags().String("dataset-group-arn", "", "The ARN of the dataset group that the filter will belong to.")
	personalize_createFilterCmd.Flags().String("filter-expression", "", "The filter expression defines which items are included or excluded from recommendations.")
	personalize_createFilterCmd.Flags().String("name", "", "The name of the filter to create.")
	personalize_createFilterCmd.Flags().String("tags", "", "A list of [tags](https://docs.aws.amazon.com/personalize/latest/dg/tagging-resources.html) to apply to the filter.")
	personalize_createFilterCmd.MarkFlagRequired("dataset-group-arn")
	personalize_createFilterCmd.MarkFlagRequired("filter-expression")
	personalize_createFilterCmd.MarkFlagRequired("name")
	personalizeCmd.AddCommand(personalize_createFilterCmd)
}
