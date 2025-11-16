package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_describeClusterV2Cmd = &cobra.Command{
	Use:   "describe-cluster-v2",
	Short: "Returns a description of the MSK cluster whose Amazon Resource Name (ARN) is specified in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_describeClusterV2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafka_describeClusterV2Cmd).Standalone()

		kafka_describeClusterV2Cmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies the cluster.")
		kafka_describeClusterV2Cmd.MarkFlagRequired("cluster-arn")
	})
	kafkaCmd.AddCommand(kafka_describeClusterV2Cmd)
}
