package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplaceCatalog_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag or list of tags from a resource (either an [entity](https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/welcome.html#catalog-api-entities) or [change set](https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/welcome.html#working-with-change-sets)).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplaceCatalog_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(marketplaceCatalog_untagResourceCmd).Standalone()

		marketplaceCatalog_untagResourceCmd.Flags().String("resource-arn", "", "Required.")
		marketplaceCatalog_untagResourceCmd.Flags().String("tag-keys", "", "Required.")
		marketplaceCatalog_untagResourceCmd.MarkFlagRequired("resource-arn")
		marketplaceCatalog_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	marketplaceCatalogCmd.AddCommand(marketplaceCatalog_untagResourceCmd)
}
