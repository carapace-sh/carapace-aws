package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_deleteVpcLinkCmd = &cobra.Command{
	Use:   "delete-vpc-link",
	Short: "Deletes an existing VpcLink of a specified identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_deleteVpcLinkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_deleteVpcLinkCmd).Standalone()

		apigateway_deleteVpcLinkCmd.Flags().String("vpc-link-id", "", "The identifier of the VpcLink.")
		apigateway_deleteVpcLinkCmd.MarkFlagRequired("vpc-link-id")
	})
	apigatewayCmd.AddCommand(apigateway_deleteVpcLinkCmd)
}
