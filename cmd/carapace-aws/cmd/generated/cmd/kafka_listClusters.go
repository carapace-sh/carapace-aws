package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_listClustersCmd = &cobra.Command{
	Use:   "list-clusters",
	Short: "Returns a list of all the MSK clusters in the current Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_listClustersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafka_listClustersCmd).Standalone()

		kafka_listClustersCmd.Flags().String("cluster-name-filter", "", "Specify a prefix of the name of the clusters that you want to list.")
		kafka_listClustersCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		kafka_listClustersCmd.Flags().String("next-token", "", "The paginated results marker.")
	})
	kafkaCmd.AddCommand(kafka_listClustersCmd)
}
