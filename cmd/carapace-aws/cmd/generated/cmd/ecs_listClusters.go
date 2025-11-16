package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_listClustersCmd = &cobra.Command{
	Use:   "list-clusters",
	Short: "Returns a list of existing clusters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_listClustersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_listClustersCmd).Standalone()

		ecs_listClustersCmd.Flags().String("max-results", "", "The maximum number of cluster results that `ListClusters` returned in paginated output.")
		ecs_listClustersCmd.Flags().String("next-token", "", "The `nextToken` value returned from a `ListClusters` request indicating that more results are available to fulfill the request and further calls are needed.")
	})
	ecsCmd.AddCommand(ecs_listClustersCmd)
}
