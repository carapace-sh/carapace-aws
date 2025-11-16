package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_deleteApiCmd = &cobra.Command{
	Use:   "delete-api",
	Short: "Deletes an Api resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_deleteApiCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_deleteApiCmd).Standalone()

		apigatewayv2_deleteApiCmd.Flags().String("api-id", "", "The API identifier.")
		apigatewayv2_deleteApiCmd.MarkFlagRequired("api-id")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_deleteApiCmd)
}
