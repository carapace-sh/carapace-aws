package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_listVersionsCmd = &cobra.Command{
	Use:   "list-versions",
	Short: "Lists all versions of OpenSearch and Elasticsearch that Amazon OpenSearch Service supports.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_listVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_listVersionsCmd).Standalone()

		opensearch_listVersionsCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
		opensearch_listVersionsCmd.Flags().String("next-token", "", "If your initial `ListVersions` operation returns a `nextToken`, you can include the returned `nextToken` in subsequent `ListVersions` operations, which returns results in the next page.")
	})
	opensearchCmd.AddCommand(opensearch_listVersionsCmd)
}
