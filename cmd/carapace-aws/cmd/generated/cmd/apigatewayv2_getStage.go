package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_getStageCmd = &cobra.Command{
	Use:   "get-stage",
	Short: "Gets a Stage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_getStageCmd).Standalone()

	apigatewayv2_getStageCmd.Flags().String("api-id", "", "The API identifier.")
	apigatewayv2_getStageCmd.Flags().String("stage-name", "", "The stage name.")
	apigatewayv2_getStageCmd.MarkFlagRequired("api-id")
	apigatewayv2_getStageCmd.MarkFlagRequired("stage-name")
	apigatewayv2Cmd.AddCommand(apigatewayv2_getStageCmd)
}
