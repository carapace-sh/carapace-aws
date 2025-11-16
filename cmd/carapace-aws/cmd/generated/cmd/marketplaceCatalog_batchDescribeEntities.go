package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplaceCatalog_batchDescribeEntitiesCmd = &cobra.Command{
	Use:   "batch-describe-entities",
	Short: "Returns metadata and content for multiple entities.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplaceCatalog_batchDescribeEntitiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(marketplaceCatalog_batchDescribeEntitiesCmd).Standalone()

		marketplaceCatalog_batchDescribeEntitiesCmd.Flags().String("entity-request-list", "", "List of entity IDs and the catalogs the entities are present in.")
		marketplaceCatalog_batchDescribeEntitiesCmd.MarkFlagRequired("entity-request-list")
	})
	marketplaceCatalogCmd.AddCommand(marketplaceCatalog_batchDescribeEntitiesCmd)
}
