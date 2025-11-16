package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplaceCatalog_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Tags a resource (either an [entity](https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/welcome.html#catalog-api-entities) or [change set](https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/welcome.html#working-with-change-sets)).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplaceCatalog_tagResourceCmd).Standalone()

	marketplaceCatalog_tagResourceCmd.Flags().String("resource-arn", "", "Required.")
	marketplaceCatalog_tagResourceCmd.Flags().String("tags", "", "Required.")
	marketplaceCatalog_tagResourceCmd.MarkFlagRequired("resource-arn")
	marketplaceCatalog_tagResourceCmd.MarkFlagRequired("tags")
	marketplaceCatalogCmd.AddCommand(marketplaceCatalog_tagResourceCmd)
}
