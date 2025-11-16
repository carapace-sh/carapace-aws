package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsql_listClustersCmd = &cobra.Command{
	Use:   "list-clusters",
	Short: "Retrieves information about a list of clusters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsql_listClustersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dsql_listClustersCmd).Standalone()

		dsql_listClustersCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
		dsql_listClustersCmd.Flags().String("next-token", "", "If your initial ListClusters operation returns a nextToken, you can include the returned nextToken in following ListClusters operations, which returns results in the next page.")
	})
	dsqlCmd.AddCommand(dsql_listClustersCmd)
}
