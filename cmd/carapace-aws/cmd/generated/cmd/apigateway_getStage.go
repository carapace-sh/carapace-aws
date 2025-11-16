package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getStageCmd = &cobra.Command{
	Use:   "get-stage",
	Short: "Gets information about a Stage resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getStageCmd).Standalone()

	apigateway_getStageCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_getStageCmd.Flags().String("stage-name", "", "The name of the Stage resource to get information about.")
	apigateway_getStageCmd.MarkFlagRequired("rest-api-id")
	apigateway_getStageCmd.MarkFlagRequired("stage-name")
	apigatewayCmd.AddCommand(apigateway_getStageCmd)
}
