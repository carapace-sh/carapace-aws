package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updateDomainConfigurationCmd = &cobra.Command{
	Use:   "update-domain-configuration",
	Short: "Updates values stored in the domain configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updateDomainConfigurationCmd).Standalone()

	iot_updateDomainConfigurationCmd.Flags().String("application-protocol", "", "An enumerated string that speciﬁes the application-layer protocol.")
	iot_updateDomainConfigurationCmd.Flags().String("authentication-type", "", "An enumerated string that speciﬁes the authentication type.")
	iot_updateDomainConfigurationCmd.Flags().String("authorizer-config", "", "An object that specifies the authorization service for a domain.")
	iot_updateDomainConfigurationCmd.Flags().String("client-certificate-config", "", "An object that speciﬁes the client certificate conﬁguration for a domain.")
	iot_updateDomainConfigurationCmd.Flags().String("domain-configuration-name", "", "The name of the domain configuration to be updated.")
	iot_updateDomainConfigurationCmd.Flags().String("domain-configuration-status", "", "The status to which the domain configuration should be updated.")
	iot_updateDomainConfigurationCmd.Flags().String("remove-authorizer-config", "", "Removes the authorization configuration from a domain.")
	iot_updateDomainConfigurationCmd.Flags().String("server-certificate-config", "", "The server certificate configuration.")
	iot_updateDomainConfigurationCmd.Flags().String("tls-config", "", "An object that specifies the TLS configuration for a domain.")
	iot_updateDomainConfigurationCmd.MarkFlagRequired("domain-configuration-name")
	iotCmd.AddCommand(iot_updateDomainConfigurationCmd)
}
