package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_batchGetCollectionCmd = &cobra.Command{
	Use:   "batch-get-collection",
	Short: "Returns attributes for one or more collections, including the collection endpoint, the OpenSearch Dashboards endpoint, and FIPS-compliant endpoints.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_batchGetCollectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearchserverless_batchGetCollectionCmd).Standalone()

		opensearchserverless_batchGetCollectionCmd.Flags().String("ids", "", "A list of collection IDs.")
		opensearchserverless_batchGetCollectionCmd.Flags().String("names", "", "A list of collection names.")
	})
	opensearchserverlessCmd.AddCommand(opensearchserverless_batchGetCollectionCmd)
}
