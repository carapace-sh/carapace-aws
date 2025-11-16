package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeGlobalClustersCmd = &cobra.Command{
	Use:   "describe-global-clusters",
	Short: "Returns information about Aurora global database clusters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeGlobalClustersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_describeGlobalClustersCmd).Standalone()

		rds_describeGlobalClustersCmd.Flags().String("filters", "", "A filter that specifies one or more global database clusters to describe.")
		rds_describeGlobalClustersCmd.Flags().String("global-cluster-identifier", "", "The user-supplied DB cluster identifier.")
		rds_describeGlobalClustersCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeGlobalClusters` request.")
		rds_describeGlobalClustersCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	})
	rdsCmd.AddCommand(rds_describeGlobalClustersCmd)
}
