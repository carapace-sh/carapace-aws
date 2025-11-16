package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_describeClustersCmd = &cobra.Command{
	Use:   "describe-clusters",
	Short: "Returns information about all provisioned clusters if no cluster identifier is specified, or about a specific cluster if a cluster name is supplied.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_describeClustersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(memorydb_describeClustersCmd).Standalone()

		memorydb_describeClustersCmd.Flags().String("cluster-name", "", "The name of the cluster.")
		memorydb_describeClustersCmd.Flags().String("max-results", "", "The maximum number of records to include in the response.")
		memorydb_describeClustersCmd.Flags().String("next-token", "", "An optional argument to pass in case the total number of records exceeds the value of MaxResults.")
		memorydb_describeClustersCmd.Flags().String("show-shard-details", "", "An optional flag that can be included in the request to retrieve information about the individual shard(s).")
	})
	memorydbCmd.AddCommand(memorydb_describeClustersCmd)
}
