package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_describeClusterCmd = &cobra.Command{
	Use:   "describe-cluster",
	Short: "Returns a description of the MSK cluster whose Amazon Resource Name (ARN) is specified in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_describeClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafka_describeClusterCmd).Standalone()

		kafka_describeClusterCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies the cluster.")
		kafka_describeClusterCmd.MarkFlagRequired("cluster-arn")
	})
	kafkaCmd.AddCommand(kafka_describeClusterCmd)
}
