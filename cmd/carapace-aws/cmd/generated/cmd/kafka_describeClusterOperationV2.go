package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_describeClusterOperationV2Cmd = &cobra.Command{
	Use:   "describe-cluster-operation-v2",
	Short: "Returns a description of the cluster operation specified by the ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_describeClusterOperationV2Cmd).Standalone()

	kafka_describeClusterOperationV2Cmd.Flags().String("cluster-operation-arn", "", "ARN of the cluster operation to describe.")
	kafka_describeClusterOperationV2Cmd.MarkFlagRequired("cluster-operation-arn")
	kafkaCmd.AddCommand(kafka_describeClusterOperationV2Cmd)
}
