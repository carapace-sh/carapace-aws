package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_listResourcesForTagOptionCmd = &cobra.Command{
	Use:   "list-resources-for-tag-option",
	Short: "Lists the resources associated with the specified TagOption.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_listResourcesForTagOptionCmd).Standalone()

	servicecatalog_listResourcesForTagOptionCmd.Flags().String("page-size", "", "The maximum number of items to return with this call.")
	servicecatalog_listResourcesForTagOptionCmd.Flags().String("page-token", "", "The page token for the next set of results.")
	servicecatalog_listResourcesForTagOptionCmd.Flags().String("resource-type", "", "The resource type.")
	servicecatalog_listResourcesForTagOptionCmd.Flags().String("tag-option-id", "", "The TagOption identifier.")
	servicecatalog_listResourcesForTagOptionCmd.MarkFlagRequired("tag-option-id")
	servicecatalogCmd.AddCommand(servicecatalog_listResourcesForTagOptionCmd)
}
