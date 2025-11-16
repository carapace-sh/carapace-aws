package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_updateReplicationInfoCmd = &cobra.Command{
	Use:   "update-replication-info",
	Short: "Updates replication info of a replicator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_updateReplicationInfoCmd).Standalone()

	kafka_updateReplicationInfoCmd.Flags().String("consumer-group-replication", "", "Updated consumer group replication information.")
	kafka_updateReplicationInfoCmd.Flags().String("current-version", "", "Current replicator version.")
	kafka_updateReplicationInfoCmd.Flags().String("replicator-arn", "", "The Amazon Resource Name (ARN) of the replicator to be updated.")
	kafka_updateReplicationInfoCmd.Flags().String("source-kafka-cluster-arn", "", "The ARN of the source Kafka cluster.")
	kafka_updateReplicationInfoCmd.Flags().String("target-kafka-cluster-arn", "", "The ARN of the target Kafka cluster.")
	kafka_updateReplicationInfoCmd.Flags().String("topic-replication", "", "Updated topic replication information.")
	kafka_updateReplicationInfoCmd.MarkFlagRequired("current-version")
	kafka_updateReplicationInfoCmd.MarkFlagRequired("replicator-arn")
	kafka_updateReplicationInfoCmd.MarkFlagRequired("source-kafka-cluster-arn")
	kafka_updateReplicationInfoCmd.MarkFlagRequired("target-kafka-cluster-arn")
	kafkaCmd.AddCommand(kafka_updateReplicationInfoCmd)
}
