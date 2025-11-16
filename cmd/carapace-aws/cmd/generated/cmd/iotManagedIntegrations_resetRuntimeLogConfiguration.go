package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_resetRuntimeLogConfigurationCmd = &cobra.Command{
	Use:   "reset-runtime-log-configuration",
	Short: "Reset a runtime log configuration for a specific managed thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_resetRuntimeLogConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_resetRuntimeLogConfigurationCmd).Standalone()

		iotManagedIntegrations_resetRuntimeLogConfigurationCmd.Flags().String("managed-thing-id", "", "The id of a managed thing.")
		iotManagedIntegrations_resetRuntimeLogConfigurationCmd.MarkFlagRequired("managed-thing-id")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_resetRuntimeLogConfigurationCmd)
}
