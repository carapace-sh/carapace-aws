package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_deleteClusterCmd = &cobra.Command{
	Use:   "delete-cluster",
	Short: "Deletes the MSK cluster specified by the Amazon Resource Name (ARN) in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_deleteClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafka_deleteClusterCmd).Standalone()

		kafka_deleteClusterCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies the cluster.")
		kafka_deleteClusterCmd.Flags().String("current-version", "", "The current version of the MSK cluster.")
		kafka_deleteClusterCmd.MarkFlagRequired("cluster-arn")
	})
	kafkaCmd.AddCommand(kafka_deleteClusterCmd)
}
