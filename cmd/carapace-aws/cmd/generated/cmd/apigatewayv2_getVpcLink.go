package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_getVpcLinkCmd = &cobra.Command{
	Use:   "get-vpc-link",
	Short: "Gets a VPC link.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_getVpcLinkCmd).Standalone()

	apigatewayv2_getVpcLinkCmd.Flags().String("vpc-link-id", "", "The ID of the VPC link.")
	apigatewayv2_getVpcLinkCmd.MarkFlagRequired("vpc-link-id")
	apigatewayv2Cmd.AddCommand(apigatewayv2_getVpcLinkCmd)
}
