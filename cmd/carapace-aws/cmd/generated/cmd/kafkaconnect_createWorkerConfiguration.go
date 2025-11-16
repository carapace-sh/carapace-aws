package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafkaconnect_createWorkerConfigurationCmd = &cobra.Command{
	Use:   "create-worker-configuration",
	Short: "Creates a worker configuration using the specified properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafkaconnect_createWorkerConfigurationCmd).Standalone()

	kafkaconnect_createWorkerConfigurationCmd.Flags().String("description", "", "A summary description of the worker configuration.")
	kafkaconnect_createWorkerConfigurationCmd.Flags().String("name", "", "The name of the worker configuration.")
	kafkaconnect_createWorkerConfigurationCmd.Flags().String("properties-file-content", "", "Base64 encoded contents of connect-distributed.properties file.")
	kafkaconnect_createWorkerConfigurationCmd.Flags().String("tags", "", "The tags you want to attach to the worker configuration.")
	kafkaconnect_createWorkerConfigurationCmd.MarkFlagRequired("name")
	kafkaconnect_createWorkerConfigurationCmd.MarkFlagRequired("properties-file-content")
	kafkaconnectCmd.AddCommand(kafkaconnect_createWorkerConfigurationCmd)
}
