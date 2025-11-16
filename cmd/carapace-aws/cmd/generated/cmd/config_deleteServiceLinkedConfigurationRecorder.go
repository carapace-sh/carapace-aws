package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_deleteServiceLinkedConfigurationRecorderCmd = &cobra.Command{
	Use:   "delete-service-linked-configuration-recorder",
	Short: "Deletes an existing service-linked configuration recorder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_deleteServiceLinkedConfigurationRecorderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_deleteServiceLinkedConfigurationRecorderCmd).Standalone()

		config_deleteServiceLinkedConfigurationRecorderCmd.Flags().String("service-principal", "", "The service principal of the Amazon Web Services service for the service-linked configuration recorder that you want to delete.")
		config_deleteServiceLinkedConfigurationRecorderCmd.MarkFlagRequired("service-principal")
	})
	configCmd.AddCommand(config_deleteServiceLinkedConfigurationRecorderCmd)
}
