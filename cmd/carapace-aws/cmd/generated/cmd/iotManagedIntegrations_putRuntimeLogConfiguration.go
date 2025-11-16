package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_putRuntimeLogConfigurationCmd = &cobra.Command{
	Use:   "put-runtime-log-configuration",
	Short: "Set the runtime log configuration for a specific managed thing or for all managed things as a group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_putRuntimeLogConfigurationCmd).Standalone()

	iotManagedIntegrations_putRuntimeLogConfigurationCmd.Flags().String("managed-thing-id", "", "The id for a managed thing.")
	iotManagedIntegrations_putRuntimeLogConfigurationCmd.Flags().String("runtime-log-configurations", "", "The runtime log configuration for a managed thing.")
	iotManagedIntegrations_putRuntimeLogConfigurationCmd.MarkFlagRequired("managed-thing-id")
	iotManagedIntegrations_putRuntimeLogConfigurationCmd.MarkFlagRequired("runtime-log-configurations")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_putRuntimeLogConfigurationCmd)
}
