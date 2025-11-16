package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_getRuntimeLogConfigurationCmd = &cobra.Command{
	Use:   "get-runtime-log-configuration",
	Short: "Get the runtime log configuration for a specific managed thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_getRuntimeLogConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_getRuntimeLogConfigurationCmd).Standalone()

		iotManagedIntegrations_getRuntimeLogConfigurationCmd.Flags().String("managed-thing-id", "", "The id for a managed thing.")
		iotManagedIntegrations_getRuntimeLogConfigurationCmd.MarkFlagRequired("managed-thing-id")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_getRuntimeLogConfigurationCmd)
}
