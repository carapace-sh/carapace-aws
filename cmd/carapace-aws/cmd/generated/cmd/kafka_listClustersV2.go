package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_listClustersV2Cmd = &cobra.Command{
	Use:   "list-clusters-v2",
	Short: "Returns a list of all the MSK clusters in the current Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_listClustersV2Cmd).Standalone()

	kafka_listClustersV2Cmd.Flags().String("cluster-name-filter", "", "Specify a prefix of the names of the clusters that you want to list.")
	kafka_listClustersV2Cmd.Flags().String("cluster-type-filter", "", "Specify either PROVISIONED or SERVERLESS.")
	kafka_listClustersV2Cmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	kafka_listClustersV2Cmd.Flags().String("next-token", "", "The paginated results marker.")
	kafkaCmd.AddCommand(kafka_listClustersV2Cmd)
}
