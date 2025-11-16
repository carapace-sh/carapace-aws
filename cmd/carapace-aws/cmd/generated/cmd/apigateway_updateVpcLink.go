package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_updateVpcLinkCmd = &cobra.Command{
	Use:   "update-vpc-link",
	Short: "Updates an existing VpcLink of a specified identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_updateVpcLinkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_updateVpcLinkCmd).Standalone()

		apigateway_updateVpcLinkCmd.Flags().String("patch-operations", "", "For more information about supported patch operations, see [Patch Operations](https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html).")
		apigateway_updateVpcLinkCmd.Flags().String("vpc-link-id", "", "The identifier of the VpcLink.")
		apigateway_updateVpcLinkCmd.MarkFlagRequired("vpc-link-id")
	})
	apigatewayCmd.AddCommand(apigateway_updateVpcLinkCmd)
}
