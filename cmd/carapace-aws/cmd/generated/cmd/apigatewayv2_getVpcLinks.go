package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_getVpcLinksCmd = &cobra.Command{
	Use:   "get-vpc-links",
	Short: "Gets a collection of VPC links.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_getVpcLinksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_getVpcLinksCmd).Standalone()

		apigatewayv2_getVpcLinksCmd.Flags().String("max-results", "", "The maximum number of elements to be returned for this resource.")
		apigatewayv2_getVpcLinksCmd.Flags().String("next-token", "", "The next page of elements from this collection.")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_getVpcLinksCmd)
}
