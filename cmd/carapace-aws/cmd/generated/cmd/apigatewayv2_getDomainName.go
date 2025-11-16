package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_getDomainNameCmd = &cobra.Command{
	Use:   "get-domain-name",
	Short: "Gets a domain name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_getDomainNameCmd).Standalone()

	apigatewayv2_getDomainNameCmd.Flags().String("domain-name", "", "The domain name.")
	apigatewayv2_getDomainNameCmd.MarkFlagRequired("domain-name")
	apigatewayv2Cmd.AddCommand(apigatewayv2_getDomainNameCmd)
}
