package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_createVpcLinkCmd = &cobra.Command{
	Use:   "create-vpc-link",
	Short: "Creates a VPC link, under the caller's account in a selected region, in an asynchronous operation that typically takes 2-4 minutes to complete and become operational.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_createVpcLinkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_createVpcLinkCmd).Standalone()

		apigateway_createVpcLinkCmd.Flags().String("description", "", "The description of the VPC link.")
		apigateway_createVpcLinkCmd.Flags().String("name", "", "The name used to label and identify the VPC link.")
		apigateway_createVpcLinkCmd.Flags().String("tags", "", "The key-value map of strings.")
		apigateway_createVpcLinkCmd.Flags().String("target-arns", "", "The ARN of the network load balancer of the VPC targeted by the VPC link.")
		apigateway_createVpcLinkCmd.MarkFlagRequired("name")
		apigateway_createVpcLinkCmd.MarkFlagRequired("target-arns")
	})
	apigatewayCmd.AddCommand(apigateway_createVpcLinkCmd)
}
