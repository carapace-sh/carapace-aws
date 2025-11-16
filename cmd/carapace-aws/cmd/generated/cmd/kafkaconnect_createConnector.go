package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafkaconnect_createConnectorCmd = &cobra.Command{
	Use:   "create-connector",
	Short: "Creates a connector using the specified properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafkaconnect_createConnectorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafkaconnect_createConnectorCmd).Standalone()

		kafkaconnect_createConnectorCmd.Flags().String("capacity", "", "Information about the capacity allocated to the connector.")
		kafkaconnect_createConnectorCmd.Flags().String("connector-configuration", "", "A map of keys to values that represent the configuration for the connector.")
		kafkaconnect_createConnectorCmd.Flags().String("connector-description", "", "A summary description of the connector.")
		kafkaconnect_createConnectorCmd.Flags().String("connector-name", "", "The name of the connector.")
		kafkaconnect_createConnectorCmd.Flags().String("kafka-cluster", "", "Specifies which Apache Kafka cluster to connect to.")
		kafkaconnect_createConnectorCmd.Flags().String("kafka-cluster-client-authentication", "", "Details of the client authentication used by the Apache Kafka cluster.")
		kafkaconnect_createConnectorCmd.Flags().String("kafka-cluster-encryption-in-transit", "", "Details of encryption in transit to the Apache Kafka cluster.")
		kafkaconnect_createConnectorCmd.Flags().String("kafka-connect-version", "", "The version of Kafka Connect.")
		kafkaconnect_createConnectorCmd.Flags().String("log-delivery", "", "Details about log delivery.")
		kafkaconnect_createConnectorCmd.Flags().String("plugins", "", "Amazon MSK Connect does not currently support specifying multiple plugins as a list.")
		kafkaconnect_createConnectorCmd.Flags().String("service-execution-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role used by the connector to access the Amazon Web Services resources that it needs.")
		kafkaconnect_createConnectorCmd.Flags().String("tags", "", "The tags you want to attach to the connector.")
		kafkaconnect_createConnectorCmd.Flags().String("worker-configuration", "", "Specifies which worker configuration to use with the connector.")
		kafkaconnect_createConnectorCmd.MarkFlagRequired("capacity")
		kafkaconnect_createConnectorCmd.MarkFlagRequired("connector-configuration")
		kafkaconnect_createConnectorCmd.MarkFlagRequired("connector-name")
		kafkaconnect_createConnectorCmd.MarkFlagRequired("kafka-cluster")
		kafkaconnect_createConnectorCmd.MarkFlagRequired("kafka-cluster-client-authentication")
		kafkaconnect_createConnectorCmd.MarkFlagRequired("kafka-cluster-encryption-in-transit")
		kafkaconnect_createConnectorCmd.MarkFlagRequired("kafka-connect-version")
		kafkaconnect_createConnectorCmd.MarkFlagRequired("plugins")
		kafkaconnect_createConnectorCmd.MarkFlagRequired("service-execution-role-arn")
	})
	kafkaconnectCmd.AddCommand(kafkaconnect_createConnectorCmd)
}
