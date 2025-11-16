package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafkaconnect_deleteWorkerConfigurationCmd = &cobra.Command{
	Use:   "delete-worker-configuration",
	Short: "Deletes the specified worker configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafkaconnect_deleteWorkerConfigurationCmd).Standalone()

	kafkaconnect_deleteWorkerConfigurationCmd.Flags().String("worker-configuration-arn", "", "The Amazon Resource Name (ARN) of the worker configuration that you want to delete.")
	kafkaconnect_deleteWorkerConfigurationCmd.MarkFlagRequired("worker-configuration-arn")
	kafkaconnectCmd.AddCommand(kafkaconnect_deleteWorkerConfigurationCmd)
}
