package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createDomainConfigurationCmd = &cobra.Command{
	Use:   "create-domain-configuration",
	Short: "Creates a domain configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createDomainConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_createDomainConfigurationCmd).Standalone()

		iot_createDomainConfigurationCmd.Flags().String("application-protocol", "", "An enumerated string that speciﬁes the application-layer protocol.")
		iot_createDomainConfigurationCmd.Flags().String("authentication-type", "", "An enumerated string that speciﬁes the authentication type.")
		iot_createDomainConfigurationCmd.Flags().String("authorizer-config", "", "An object that specifies the authorization service for a domain.")
		iot_createDomainConfigurationCmd.Flags().String("client-certificate-config", "", "An object that speciﬁes the client certificate conﬁguration for a domain.")
		iot_createDomainConfigurationCmd.Flags().String("domain-configuration-name", "", "The name of the domain configuration.")
		iot_createDomainConfigurationCmd.Flags().String("domain-name", "", "The name of the domain.")
		iot_createDomainConfigurationCmd.Flags().String("server-certificate-arns", "", "The ARNs of the certificates that IoT passes to the device during the TLS handshake.")
		iot_createDomainConfigurationCmd.Flags().String("server-certificate-config", "", "The server certificate configuration.")
		iot_createDomainConfigurationCmd.Flags().String("service-type", "", "The type of service delivered by the endpoint.")
		iot_createDomainConfigurationCmd.Flags().String("tags", "", "Metadata which can be used to manage the domain configuration.")
		iot_createDomainConfigurationCmd.Flags().String("tls-config", "", "An object that specifies the TLS configuration for a domain.")
		iot_createDomainConfigurationCmd.Flags().String("validation-certificate-arn", "", "The certificate used to validate the server certificate and prove domain name ownership.")
		iot_createDomainConfigurationCmd.MarkFlagRequired("domain-configuration-name")
	})
	iotCmd.AddCommand(iot_createDomainConfigurationCmd)
}
