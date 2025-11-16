package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_describeClusterOperationCmd = &cobra.Command{
	Use:   "describe-cluster-operation",
	Short: "Returns a description of the cluster operation specified by the ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_describeClusterOperationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafka_describeClusterOperationCmd).Standalone()

		kafka_describeClusterOperationCmd.Flags().String("cluster-operation-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies the MSK cluster operation.")
		kafka_describeClusterOperationCmd.MarkFlagRequired("cluster-operation-arn")
	})
	kafkaCmd.AddCommand(kafka_describeClusterOperationCmd)
}
