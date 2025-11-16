package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_createReplicatorCmd = &cobra.Command{
	Use:   "create-replicator",
	Short: "Creates the replicator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_createReplicatorCmd).Standalone()

	kafka_createReplicatorCmd.Flags().String("description", "", "A summary description of the replicator.")
	kafka_createReplicatorCmd.Flags().String("kafka-clusters", "", "Kafka Clusters to use in setting up sources / targets for replication.")
	kafka_createReplicatorCmd.Flags().String("replication-info-list", "", "A list of replication configurations, where each configuration targets a given source cluster to target cluster replication flow.")
	kafka_createReplicatorCmd.Flags().String("replicator-name", "", "The name of the replicator.")
	kafka_createReplicatorCmd.Flags().String("service-execution-role-arn", "", "The ARN of the IAM role used by the replicator to access resources in the customer's account (e.g source and target clusters)")
	kafka_createReplicatorCmd.Flags().String("tags", "", "List of tags to attach to created Replicator.")
	kafka_createReplicatorCmd.MarkFlagRequired("kafka-clusters")
	kafka_createReplicatorCmd.MarkFlagRequired("replication-info-list")
	kafka_createReplicatorCmd.MarkFlagRequired("replicator-name")
	kafka_createReplicatorCmd.MarkFlagRequired("service-execution-role-arn")
	kafkaCmd.AddCommand(kafka_createReplicatorCmd)
}
