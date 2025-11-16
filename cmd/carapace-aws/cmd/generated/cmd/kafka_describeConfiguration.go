package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_describeConfigurationCmd = &cobra.Command{
	Use:   "describe-configuration",
	Short: "Returns a description of this MSK configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_describeConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafka_describeConfigurationCmd).Standalone()

		kafka_describeConfigurationCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) that uniquely identifies an MSK configuration and all of its revisions.")
		kafka_describeConfigurationCmd.MarkFlagRequired("arn")
	})
	kafkaCmd.AddCommand(kafka_describeConfigurationCmd)
}
