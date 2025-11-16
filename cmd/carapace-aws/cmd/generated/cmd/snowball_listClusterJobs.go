package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowball_listClusterJobsCmd = &cobra.Command{
	Use:   "list-cluster-jobs",
	Short: "Returns an array of `JobListEntry` objects of the specified length.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowball_listClusterJobsCmd).Standalone()

	snowball_listClusterJobsCmd.Flags().String("cluster-id", "", "The 39-character ID for the cluster that you want to list, for example `CID123e4567-e89b-12d3-a456-426655440000`.")
	snowball_listClusterJobsCmd.Flags().String("max-results", "", "The number of `JobListEntry` objects to return.")
	snowball_listClusterJobsCmd.Flags().String("next-token", "", "HTTP requests are stateless.")
	snowball_listClusterJobsCmd.MarkFlagRequired("cluster-id")
	snowballCmd.AddCommand(snowball_listClusterJobsCmd)
}
