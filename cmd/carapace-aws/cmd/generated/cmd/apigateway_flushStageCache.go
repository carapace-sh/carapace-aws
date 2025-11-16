package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_flushStageCacheCmd = &cobra.Command{
	Use:   "flush-stage-cache",
	Short: "Flushes a stage's cache.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_flushStageCacheCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_flushStageCacheCmd).Standalone()

		apigateway_flushStageCacheCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_flushStageCacheCmd.Flags().String("stage-name", "", "The name of the stage to flush its cache.")
		apigateway_flushStageCacheCmd.MarkFlagRequired("rest-api-id")
		apigateway_flushStageCacheCmd.MarkFlagRequired("stage-name")
	})
	apigatewayCmd.AddCommand(apigateway_flushStageCacheCmd)
}
