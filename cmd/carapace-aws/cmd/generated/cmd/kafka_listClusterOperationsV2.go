package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_listClusterOperationsV2Cmd = &cobra.Command{
	Use:   "list-cluster-operations-v2",
	Short: "Returns a list of all the operations that have been performed on the specified MSK cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_listClusterOperationsV2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafka_listClusterOperationsV2Cmd).Standalone()

		kafka_listClusterOperationsV2Cmd.Flags().String("cluster-arn", "", "The arn of the cluster whose operations are being requested.")
		kafka_listClusterOperationsV2Cmd.Flags().String("max-results", "", "The maxResults of the query.")
		kafka_listClusterOperationsV2Cmd.Flags().String("next-token", "", "The nextToken of the query.")
		kafka_listClusterOperationsV2Cmd.MarkFlagRequired("cluster-arn")
	})
	kafkaCmd.AddCommand(kafka_listClusterOperationsV2Cmd)
}
