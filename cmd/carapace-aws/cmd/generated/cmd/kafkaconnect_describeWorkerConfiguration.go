package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafkaconnect_describeWorkerConfigurationCmd = &cobra.Command{
	Use:   "describe-worker-configuration",
	Short: "Returns information about a worker configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafkaconnect_describeWorkerConfigurationCmd).Standalone()

	kafkaconnect_describeWorkerConfigurationCmd.Flags().String("worker-configuration-arn", "", "The Amazon Resource Name (ARN) of the worker configuration that you want to get information about.")
	kafkaconnect_describeWorkerConfigurationCmd.MarkFlagRequired("worker-configuration-arn")
	kafkaconnectCmd.AddCommand(kafkaconnect_describeWorkerConfigurationCmd)
}
