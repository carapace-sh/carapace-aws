package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getVpcLinkCmd = &cobra.Command{
	Use:   "get-vpc-link",
	Short: "Gets a specified VPC link under the caller's account in a region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getVpcLinkCmd).Standalone()

	apigateway_getVpcLinkCmd.Flags().String("vpc-link-id", "", "The identifier of the VpcLink.")
	apigateway_getVpcLinkCmd.MarkFlagRequired("vpc-link-id")
	apigatewayCmd.AddCommand(apigateway_getVpcLinkCmd)
}
