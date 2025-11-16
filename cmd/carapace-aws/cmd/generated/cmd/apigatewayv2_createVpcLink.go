package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_createVpcLinkCmd = &cobra.Command{
	Use:   "create-vpc-link",
	Short: "Creates a VPC link.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_createVpcLinkCmd).Standalone()

	apigatewayv2_createVpcLinkCmd.Flags().String("name", "", "The name of the VPC link.")
	apigatewayv2_createVpcLinkCmd.Flags().String("security-group-ids", "", "A list of security group IDs for the VPC link.")
	apigatewayv2_createVpcLinkCmd.Flags().String("subnet-ids", "", "A list of subnet IDs to include in the VPC link.")
	apigatewayv2_createVpcLinkCmd.Flags().String("tags", "", "A list of tags.")
	apigatewayv2_createVpcLinkCmd.MarkFlagRequired("name")
	apigatewayv2_createVpcLinkCmd.MarkFlagRequired("subnet-ids")
	apigatewayv2Cmd.AddCommand(apigatewayv2_createVpcLinkCmd)
}
