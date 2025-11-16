package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_describeConfigurationRevisionCmd = &cobra.Command{
	Use:   "describe-configuration-revision",
	Short: "Returns a description of this revision of the configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_describeConfigurationRevisionCmd).Standalone()

	kafka_describeConfigurationRevisionCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) that uniquely identifies an MSK configuration and all of its revisions.")
	kafka_describeConfigurationRevisionCmd.Flags().String("revision", "", "A string that uniquely identifies a revision of an MSK configuration.")
	kafka_describeConfigurationRevisionCmd.MarkFlagRequired("arn")
	kafka_describeConfigurationRevisionCmd.MarkFlagRequired("revision")
	kafkaCmd.AddCommand(kafka_describeConfigurationRevisionCmd)
}
