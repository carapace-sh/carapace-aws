package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_updateBrokerCountCmd = &cobra.Command{
	Use:   "update-broker-count",
	Short: "Updates the number of broker nodes in the cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_updateBrokerCountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafka_updateBrokerCountCmd).Standalone()

		kafka_updateBrokerCountCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies the cluster.")
		kafka_updateBrokerCountCmd.Flags().String("current-version", "", "The version of cluster to update from.")
		kafka_updateBrokerCountCmd.Flags().String("target-number-of-broker-nodes", "", "The number of broker nodes that you want the cluster to have after this operation completes successfully.")
		kafka_updateBrokerCountCmd.MarkFlagRequired("cluster-arn")
		kafka_updateBrokerCountCmd.MarkFlagRequired("current-version")
		kafka_updateBrokerCountCmd.MarkFlagRequired("target-number-of-broker-nodes")
	})
	kafkaCmd.AddCommand(kafka_updateBrokerCountCmd)
}
