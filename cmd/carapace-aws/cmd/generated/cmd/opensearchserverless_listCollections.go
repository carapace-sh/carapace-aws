package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_listCollectionsCmd = &cobra.Command{
	Use:   "list-collections",
	Short: "Lists all OpenSearch Serverless collections.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_listCollectionsCmd).Standalone()

	opensearchserverless_listCollectionsCmd.Flags().String("collection-filters", "", "A list of filter names and values that you can use for requests.")
	opensearchserverless_listCollectionsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	opensearchserverless_listCollectionsCmd.Flags().String("next-token", "", "If your initial `ListCollections` operation returns a `nextToken`, you can include the returned `nextToken` in subsequent `ListCollections` operations, which returns results in the next page.")
	opensearchserverlessCmd.AddCommand(opensearchserverless_listCollectionsCmd)
}
