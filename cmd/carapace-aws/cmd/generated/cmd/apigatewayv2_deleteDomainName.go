package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_deleteDomainNameCmd = &cobra.Command{
	Use:   "delete-domain-name",
	Short: "Deletes a domain name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_deleteDomainNameCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_deleteDomainNameCmd).Standalone()

		apigatewayv2_deleteDomainNameCmd.Flags().String("domain-name", "", "The domain name.")
		apigatewayv2_deleteDomainNameCmd.MarkFlagRequired("domain-name")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_deleteDomainNameCmd)
}
