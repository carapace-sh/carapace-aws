package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_deleteStageCmd = &cobra.Command{
	Use:   "delete-stage",
	Short: "Deletes a Stage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_deleteStageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_deleteStageCmd).Standalone()

		apigatewayv2_deleteStageCmd.Flags().String("api-id", "", "The API identifier.")
		apigatewayv2_deleteStageCmd.Flags().String("stage-name", "", "The stage name.")
		apigatewayv2_deleteStageCmd.MarkFlagRequired("api-id")
		apigatewayv2_deleteStageCmd.MarkFlagRequired("stage-name")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_deleteStageCmd)
}
