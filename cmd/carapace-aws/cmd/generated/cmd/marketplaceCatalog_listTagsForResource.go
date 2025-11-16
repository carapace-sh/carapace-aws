package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplaceCatalog_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all tags that have been added to a resource (either an [entity](https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/welcome.html#catalog-api-entities) or [change set](https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/welcome.html#working-with-change-sets)).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplaceCatalog_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(marketplaceCatalog_listTagsForResourceCmd).Standalone()

		marketplaceCatalog_listTagsForResourceCmd.Flags().String("resource-arn", "", "Required.")
		marketplaceCatalog_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	marketplaceCatalogCmd.AddCommand(marketplaceCatalog_listTagsForResourceCmd)
}
