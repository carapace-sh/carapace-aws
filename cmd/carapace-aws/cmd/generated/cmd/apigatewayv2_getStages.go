package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_getStagesCmd = &cobra.Command{
	Use:   "get-stages",
	Short: "Gets the Stages for an API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_getStagesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_getStagesCmd).Standalone()

		apigatewayv2_getStagesCmd.Flags().String("api-id", "", "The API identifier.")
		apigatewayv2_getStagesCmd.Flags().String("max-results", "", "The maximum number of elements to be returned for this resource.")
		apigatewayv2_getStagesCmd.Flags().String("next-token", "", "The next page of elements from this collection.")
		apigatewayv2_getStagesCmd.MarkFlagRequired("api-id")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_getStagesCmd)
}
