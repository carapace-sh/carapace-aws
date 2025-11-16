package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_deleteConfigurationCmd = &cobra.Command{
	Use:   "delete-configuration",
	Short: "Deletes an MSK Configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_deleteConfigurationCmd).Standalone()

	kafka_deleteConfigurationCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) that uniquely identifies an MSK configuration.")
	kafka_deleteConfigurationCmd.MarkFlagRequired("arn")
	kafkaCmd.AddCommand(kafka_deleteConfigurationCmd)
}
