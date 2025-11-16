package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getAuthorizersCmd = &cobra.Command{
	Use:   "get-authorizers",
	Short: "Describe an existing Authorizers resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getAuthorizersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_getAuthorizersCmd).Standalone()

		apigateway_getAuthorizersCmd.Flags().String("limit", "", "The maximum number of returned results per page.")
		apigateway_getAuthorizersCmd.Flags().String("position", "", "The current pagination position in the paged result set.")
		apigateway_getAuthorizersCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_getAuthorizersCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_getAuthorizersCmd)
}
