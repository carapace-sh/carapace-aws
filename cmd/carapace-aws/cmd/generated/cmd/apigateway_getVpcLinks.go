package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getVpcLinksCmd = &cobra.Command{
	Use:   "get-vpc-links",
	Short: "Gets the VpcLinks collection under the caller's account in a selected region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getVpcLinksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_getVpcLinksCmd).Standalone()

		apigateway_getVpcLinksCmd.Flags().String("limit", "", "The maximum number of returned results per page.")
		apigateway_getVpcLinksCmd.Flags().String("position", "", "The current pagination position in the paged result set.")
	})
	apigatewayCmd.AddCommand(apigateway_getVpcLinksCmd)
}
