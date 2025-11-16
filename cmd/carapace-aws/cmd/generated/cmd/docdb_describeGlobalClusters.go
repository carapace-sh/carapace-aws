package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_describeGlobalClustersCmd = &cobra.Command{
	Use:   "describe-global-clusters",
	Short: "Returns information about Amazon DocumentDB global clusters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_describeGlobalClustersCmd).Standalone()

	docdb_describeGlobalClustersCmd.Flags().String("filters", "", "A filter that specifies one or more global DB clusters to describe.")
	docdb_describeGlobalClustersCmd.Flags().String("global-cluster-identifier", "", "The user-supplied cluster identifier.")
	docdb_describeGlobalClustersCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeGlobalClusters` request.")
	docdb_describeGlobalClustersCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	docdbCmd.AddCommand(docdb_describeGlobalClustersCmd)
}
