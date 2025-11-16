package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_updateDomainNameCmd = &cobra.Command{
	Use:   "update-domain-name",
	Short: "Updates a domain name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_updateDomainNameCmd).Standalone()

	apigatewayv2_updateDomainNameCmd.Flags().String("domain-name", "", "The domain name.")
	apigatewayv2_updateDomainNameCmd.Flags().String("domain-name-configurations", "", "The domain name configurations.")
	apigatewayv2_updateDomainNameCmd.Flags().String("mutual-tls-authentication", "", "The mutual TLS authentication configuration for a custom domain name.")
	apigatewayv2_updateDomainNameCmd.Flags().String("routing-mode", "", "The routing mode.")
	apigatewayv2_updateDomainNameCmd.MarkFlagRequired("domain-name")
	apigatewayv2Cmd.AddCommand(apigatewayv2_updateDomainNameCmd)
}
