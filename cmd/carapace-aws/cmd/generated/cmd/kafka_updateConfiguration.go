package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_updateConfigurationCmd = &cobra.Command{
	Use:   "update-configuration",
	Short: "Updates an MSK configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_updateConfigurationCmd).Standalone()

	kafka_updateConfigurationCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the configuration.")
	kafka_updateConfigurationCmd.Flags().String("description", "", "The description of the configuration revision.")
	kafka_updateConfigurationCmd.Flags().String("server-properties", "", "Contents of the server.properties file.")
	kafka_updateConfigurationCmd.MarkFlagRequired("arn")
	kafka_updateConfigurationCmd.MarkFlagRequired("server-properties")
	kafkaCmd.AddCommand(kafka_updateConfigurationCmd)
}
