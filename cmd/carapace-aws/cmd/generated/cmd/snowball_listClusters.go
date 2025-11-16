package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowball_listClustersCmd = &cobra.Command{
	Use:   "list-clusters",
	Short: "Returns an array of `ClusterListEntry` objects of the specified length.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowball_listClustersCmd).Standalone()

	snowball_listClustersCmd.Flags().String("max-results", "", "The number of `ClusterListEntry` objects to return.")
	snowball_listClustersCmd.Flags().String("next-token", "", "HTTP requests are stateless.")
	snowballCmd.AddCommand(snowball_listClustersCmd)
}
