package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_listKxClustersCmd = &cobra.Command{
	Use:   "list-kx-clusters",
	Short: "Returns a list of clusters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_listKxClustersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspace_listKxClustersCmd).Standalone()

		finspace_listKxClustersCmd.Flags().String("cluster-type", "", "Specifies the type of KDB database that is being created.")
		finspace_listKxClustersCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment.")
		finspace_listKxClustersCmd.Flags().String("max-results", "", "The maximum number of results to return in this request.")
		finspace_listKxClustersCmd.Flags().String("next-token", "", "A token that indicates where a results page should begin.")
		finspace_listKxClustersCmd.MarkFlagRequired("environment-id")
	})
	finspaceCmd.AddCommand(finspace_listKxClustersCmd)
}
