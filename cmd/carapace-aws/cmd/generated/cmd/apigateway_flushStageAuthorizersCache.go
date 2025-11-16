package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_flushStageAuthorizersCacheCmd = &cobra.Command{
	Use:   "flush-stage-authorizers-cache",
	Short: "Flushes all authorizer cache entries on a stage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_flushStageAuthorizersCacheCmd).Standalone()

	apigateway_flushStageAuthorizersCacheCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_flushStageAuthorizersCacheCmd.Flags().String("stage-name", "", "The name of the stage to flush.")
	apigateway_flushStageAuthorizersCacheCmd.MarkFlagRequired("rest-api-id")
	apigateway_flushStageAuthorizersCacheCmd.MarkFlagRequired("stage-name")
	apigatewayCmd.AddCommand(apigateway_flushStageAuthorizersCacheCmd)
}
