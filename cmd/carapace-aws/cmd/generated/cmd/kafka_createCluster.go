package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_createClusterCmd = &cobra.Command{
	Use:   "create-cluster",
	Short: "Creates a new MSK cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_createClusterCmd).Standalone()

	kafka_createClusterCmd.Flags().String("broker-node-group-info", "", "Information about the broker nodes in the cluster.")
	kafka_createClusterCmd.Flags().String("client-authentication", "", "Includes all client authentication related information.")
	kafka_createClusterCmd.Flags().String("cluster-name", "", "The name of the cluster.")
	kafka_createClusterCmd.Flags().String("configuration-info", "", "Represents the configuration that you want MSK to use for the brokers in a cluster.")
	kafka_createClusterCmd.Flags().String("encryption-info", "", "Includes all encryption-related information.")
	kafka_createClusterCmd.Flags().String("enhanced-monitoring", "", "Specifies the level of monitoring for the MSK cluster.")
	kafka_createClusterCmd.Flags().String("kafka-version", "", "The version of Apache Kafka.")
	kafka_createClusterCmd.Flags().String("logging-info", "", "")
	kafka_createClusterCmd.Flags().String("number-of-broker-nodes", "", "The number of broker nodes in the cluster.")
	kafka_createClusterCmd.Flags().String("open-monitoring", "", "The settings for open monitoring.")
	kafka_createClusterCmd.Flags().String("rebalancing", "", "Specifies if intelligent rebalancing should be turned on for the new MSK Provisioned cluster with Express brokers.")
	kafka_createClusterCmd.Flags().String("storage-mode", "", "This controls storage mode for supported storage tiers.")
	kafka_createClusterCmd.Flags().String("tags", "", "Create tags when creating the cluster.")
	kafka_createClusterCmd.MarkFlagRequired("broker-node-group-info")
	kafka_createClusterCmd.MarkFlagRequired("cluster-name")
	kafka_createClusterCmd.MarkFlagRequired("kafka-version")
	kafka_createClusterCmd.MarkFlagRequired("number-of-broker-nodes")
	kafkaCmd.AddCommand(kafka_createClusterCmd)
}
