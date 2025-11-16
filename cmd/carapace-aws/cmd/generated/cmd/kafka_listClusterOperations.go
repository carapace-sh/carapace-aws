package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_listClusterOperationsCmd = &cobra.Command{
	Use:   "list-cluster-operations",
	Short: "Returns a list of all the operations that have been performed on the specified MSK cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_listClusterOperationsCmd).Standalone()

	kafka_listClusterOperationsCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies the cluster.")
	kafka_listClusterOperationsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	kafka_listClusterOperationsCmd.Flags().String("next-token", "", "The paginated results marker.")
	kafka_listClusterOperationsCmd.MarkFlagRequired("cluster-arn")
	kafkaCmd.AddCommand(kafka_listClusterOperationsCmd)
}
