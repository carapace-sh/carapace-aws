package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_deleteStageCmd = &cobra.Command{
	Use:   "delete-stage",
	Short: "Deletes a Stage resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_deleteStageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_deleteStageCmd).Standalone()

		apigateway_deleteStageCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_deleteStageCmd.Flags().String("stage-name", "", "The name of the Stage resource to delete.")
		apigateway_deleteStageCmd.MarkFlagRequired("rest-api-id")
		apigateway_deleteStageCmd.MarkFlagRequired("stage-name")
	})
	apigatewayCmd.AddCommand(apigateway_deleteStageCmd)
}
