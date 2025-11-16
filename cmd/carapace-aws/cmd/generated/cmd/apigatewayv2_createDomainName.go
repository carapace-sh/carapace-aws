package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_createDomainNameCmd = &cobra.Command{
	Use:   "create-domain-name",
	Short: "Creates a domain name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_createDomainNameCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_createDomainNameCmd).Standalone()

		apigatewayv2_createDomainNameCmd.Flags().String("domain-name", "", "The domain name.")
		apigatewayv2_createDomainNameCmd.Flags().String("domain-name-configurations", "", "The domain name configurations.")
		apigatewayv2_createDomainNameCmd.Flags().String("mutual-tls-authentication", "", "The mutual TLS authentication configuration for a custom domain name.")
		apigatewayv2_createDomainNameCmd.Flags().String("routing-mode", "", "The routing mode.")
		apigatewayv2_createDomainNameCmd.Flags().String("tags", "", "The collection of tags associated with a domain name.")
		apigatewayv2_createDomainNameCmd.MarkFlagRequired("domain-name")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_createDomainNameCmd)
}
