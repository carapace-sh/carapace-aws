package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getRestApiCmd = &cobra.Command{
	Use:   "get-rest-api",
	Short: "Lists the RestApi resource in the collection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getRestApiCmd).Standalone()

	apigateway_getRestApiCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_getRestApiCmd.MarkFlagRequired("rest-api-id")
	apigatewayCmd.AddCommand(apigateway_getRestApiCmd)
}
