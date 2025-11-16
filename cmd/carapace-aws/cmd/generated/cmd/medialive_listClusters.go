package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_listClustersCmd = &cobra.Command{
	Use:   "list-clusters",
	Short: "Retrieve the list of Clusters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_listClustersCmd).Standalone()

	medialive_listClustersCmd.Flags().String("max-results", "", "The maximum number of items to return.")
	medialive_listClustersCmd.Flags().String("next-token", "", "The token to retrieve the next page of results.")
	medialiveCmd.AddCommand(medialive_listClustersCmd)
}
