package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_createConfigurationCmd = &cobra.Command{
	Use:   "create-configuration",
	Short: "Creates a new MSK configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_createConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafka_createConfigurationCmd).Standalone()

		kafka_createConfigurationCmd.Flags().String("description", "", "The description of the configuration.")
		kafka_createConfigurationCmd.Flags().String("kafka-versions", "", "The versions of Apache Kafka with which you can use this MSK configuration.")
		kafka_createConfigurationCmd.Flags().String("name", "", "The name of the configuration.")
		kafka_createConfigurationCmd.Flags().String("server-properties", "", "Contents of the server.properties file.")
		kafka_createConfigurationCmd.MarkFlagRequired("name")
		kafka_createConfigurationCmd.MarkFlagRequired("server-properties")
	})
	kafkaCmd.AddCommand(kafka_createConfigurationCmd)
}
