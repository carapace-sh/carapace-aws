package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_updateVpcLinkCmd = &cobra.Command{
	Use:   "update-vpc-link",
	Short: "Updates a VPC link.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_updateVpcLinkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_updateVpcLinkCmd).Standalone()

		apigatewayv2_updateVpcLinkCmd.Flags().String("name", "", "The name of the VPC link.")
		apigatewayv2_updateVpcLinkCmd.Flags().String("vpc-link-id", "", "The ID of the VPC link.")
		apigatewayv2_updateVpcLinkCmd.MarkFlagRequired("vpc-link-id")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_updateVpcLinkCmd)
}
