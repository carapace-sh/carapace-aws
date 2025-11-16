package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_resetAuthorizersCacheCmd = &cobra.Command{
	Use:   "reset-authorizers-cache",
	Short: "Resets all authorizer cache entries on a stage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_resetAuthorizersCacheCmd).Standalone()

	apigatewayv2_resetAuthorizersCacheCmd.Flags().String("api-id", "", "The API identifier.")
	apigatewayv2_resetAuthorizersCacheCmd.Flags().String("stage-name", "", "The stage name.")
	apigatewayv2_resetAuthorizersCacheCmd.MarkFlagRequired("api-id")
	apigatewayv2_resetAuthorizersCacheCmd.MarkFlagRequired("stage-name")
	apigatewayv2Cmd.AddCommand(apigatewayv2_resetAuthorizersCacheCmd)
}
