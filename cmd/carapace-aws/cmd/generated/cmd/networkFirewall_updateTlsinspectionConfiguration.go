package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewall_updateTlsinspectionConfigurationCmd = &cobra.Command{
	Use:   "update-tlsinspection-configuration",
	Short: "Updates the TLS inspection configuration settings for the specified TLS inspection configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewall_updateTlsinspectionConfigurationCmd).Standalone()

	networkFirewall_updateTlsinspectionConfigurationCmd.Flags().String("description", "", "A description of the TLS inspection configuration.")
	networkFirewall_updateTlsinspectionConfigurationCmd.Flags().String("encryption-configuration", "", "A complex type that contains the Amazon Web Services KMS encryption configuration settings for your TLS inspection configuration.")
	networkFirewall_updateTlsinspectionConfigurationCmd.Flags().String("tlsinspection-configuration", "", "The object that defines a TLS inspection configuration.")
	networkFirewall_updateTlsinspectionConfigurationCmd.Flags().String("tlsinspection-configuration-arn", "", "The Amazon Resource Name (ARN) of the TLS inspection configuration.")
	networkFirewall_updateTlsinspectionConfigurationCmd.Flags().String("tlsinspection-configuration-name", "", "The descriptive name of the TLS inspection configuration.")
	networkFirewall_updateTlsinspectionConfigurationCmd.Flags().String("update-token", "", "A token used for optimistic locking.")
	networkFirewall_updateTlsinspectionConfigurationCmd.MarkFlagRequired("tlsinspection-configuration")
	networkFirewall_updateTlsinspectionConfigurationCmd.MarkFlagRequired("update-token")
	networkFirewallCmd.AddCommand(networkFirewall_updateTlsinspectionConfigurationCmd)
}
