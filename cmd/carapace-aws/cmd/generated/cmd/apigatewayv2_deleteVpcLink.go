package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_deleteVpcLinkCmd = &cobra.Command{
	Use:   "delete-vpc-link",
	Short: "Deletes a VPC link.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_deleteVpcLinkCmd).Standalone()

	apigatewayv2_deleteVpcLinkCmd.Flags().String("vpc-link-id", "", "The ID of the VPC link.")
	apigatewayv2_deleteVpcLinkCmd.MarkFlagRequired("vpc-link-id")
	apigatewayv2Cmd.AddCommand(apigatewayv2_deleteVpcLinkCmd)
}
