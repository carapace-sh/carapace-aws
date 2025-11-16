package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_putServiceLinkedConfigurationRecorderCmd = &cobra.Command{
	Use:   "put-service-linked-configuration-recorder",
	Short: "Creates a service-linked configuration recorder that is linked to a specific Amazon Web Services service based on the `ServicePrincipal` you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_putServiceLinkedConfigurationRecorderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_putServiceLinkedConfigurationRecorderCmd).Standalone()

		config_putServiceLinkedConfigurationRecorderCmd.Flags().String("service-principal", "", "The service principal of the Amazon Web Services service for the service-linked configuration recorder that you want to create.")
		config_putServiceLinkedConfigurationRecorderCmd.Flags().String("tags", "", "The tags for a service-linked configuration recorder.")
		config_putServiceLinkedConfigurationRecorderCmd.MarkFlagRequired("service-principal")
	})
	configCmd.AddCommand(config_putServiceLinkedConfigurationRecorderCmd)
}
