package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_createDomainNameCmd = &cobra.Command{
	Use:   "create-domain-name",
	Short: "Creates a new domain name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_createDomainNameCmd).Standalone()

	apigateway_createDomainNameCmd.Flags().String("certificate-arn", "", "The reference to an Amazon Web Services-managed certificate that will be used by edge-optimized endpoint or private endpoint for this domain name.")
	apigateway_createDomainNameCmd.Flags().String("certificate-body", "", "\\[Deprecated] The body of the server certificate that will be used by edge-optimized endpoint or private endpoint for this domain name provided by your certificate authority.")
	apigateway_createDomainNameCmd.Flags().String("certificate-chain", "", "\\[Deprecated] The intermediate certificates and optionally the root certificate, one after the other without any blank lines, used by an edge-optimized endpoint for this domain name.")
	apigateway_createDomainNameCmd.Flags().String("certificate-name", "", "The user-friendly name of the certificate that will be used by edge-optimized endpoint or private endpoint for this domain name.")
	apigateway_createDomainNameCmd.Flags().String("certificate-private-key", "", "\\[Deprecated] Your edge-optimized endpoint's domain name certificate's private key.")
	apigateway_createDomainNameCmd.Flags().String("domain-name", "", "The name of the DomainName resource.")
	apigateway_createDomainNameCmd.Flags().String("endpoint-configuration", "", "The endpoint configuration of this DomainName showing the endpoint types and IP address types of the domain name.")
	apigateway_createDomainNameCmd.Flags().String("mutual-tls-authentication", "", "")
	apigateway_createDomainNameCmd.Flags().String("ownership-verification-certificate-arn", "", "The ARN of the public certificate issued by ACM to validate ownership of your custom domain.")
	apigateway_createDomainNameCmd.Flags().String("policy", "", "A stringified JSON policy document that applies to the `execute-api` service for this DomainName regardless of the caller and Method configuration.")
	apigateway_createDomainNameCmd.Flags().String("regional-certificate-arn", "", "The reference to an Amazon Web Services-managed certificate that will be used by regional endpoint for this domain name.")
	apigateway_createDomainNameCmd.Flags().String("regional-certificate-name", "", "The user-friendly name of the certificate that will be used by regional endpoint for this domain name.")
	apigateway_createDomainNameCmd.Flags().String("routing-mode", "", "The routing mode for this domain name.")
	apigateway_createDomainNameCmd.Flags().String("security-policy", "", "The Transport Layer Security (TLS) version + cipher suite for this DomainName.")
	apigateway_createDomainNameCmd.Flags().String("tags", "", "The key-value map of strings.")
	apigateway_createDomainNameCmd.MarkFlagRequired("domain-name")
	apigatewayCmd.AddCommand(apigateway_createDomainNameCmd)
}
