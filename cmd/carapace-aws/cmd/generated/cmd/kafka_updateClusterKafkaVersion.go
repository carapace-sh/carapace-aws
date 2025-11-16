package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_updateClusterKafkaVersionCmd = &cobra.Command{
	Use:   "update-cluster-kafka-version",
	Short: "Updates the Apache Kafka version for the cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_updateClusterKafkaVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafka_updateClusterKafkaVersionCmd).Standalone()

		kafka_updateClusterKafkaVersionCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) of the cluster to be updated.")
		kafka_updateClusterKafkaVersionCmd.Flags().String("configuration-info", "", "The custom configuration that should be applied on the new version of cluster.")
		kafka_updateClusterKafkaVersionCmd.Flags().String("current-version", "", "Current cluster version.")
		kafka_updateClusterKafkaVersionCmd.Flags().String("target-kafka-version", "", "Target Kafka version.")
		kafka_updateClusterKafkaVersionCmd.MarkFlagRequired("cluster-arn")
		kafka_updateClusterKafkaVersionCmd.MarkFlagRequired("current-version")
		kafka_updateClusterKafkaVersionCmd.MarkFlagRequired("target-kafka-version")
	})
	kafkaCmd.AddCommand(kafka_updateClusterKafkaVersionCmd)
}
