package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_describeGlobalClustersCmd = &cobra.Command{
	Use:   "describe-global-clusters",
	Short: "Returns information about Neptune global database clusters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_describeGlobalClustersCmd).Standalone()

	neptune_describeGlobalClustersCmd.Flags().String("global-cluster-identifier", "", "The user-supplied DB cluster identifier.")
	neptune_describeGlobalClustersCmd.Flags().String("marker", "", "(*Optional*) A pagination token returned by a previous call to `DescribeGlobalClusters`.")
	neptune_describeGlobalClustersCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	neptuneCmd.AddCommand(neptune_describeGlobalClustersCmd)
}
