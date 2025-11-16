package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_getDomainNamesCmd = &cobra.Command{
	Use:   "get-domain-names",
	Short: "Gets the domain names for an AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_getDomainNamesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_getDomainNamesCmd).Standalone()

		apigatewayv2_getDomainNamesCmd.Flags().String("max-results", "", "The maximum number of elements to be returned for this resource.")
		apigatewayv2_getDomainNamesCmd.Flags().String("next-token", "", "The next page of elements from this collection.")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_getDomainNamesCmd)
}
