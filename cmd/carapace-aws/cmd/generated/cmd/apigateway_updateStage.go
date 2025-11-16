package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_updateStageCmd = &cobra.Command{
	Use:   "update-stage",
	Short: "Changes information about a Stage resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_updateStageCmd).Standalone()

	apigateway_updateStageCmd.Flags().String("patch-operations", "", "For more information about supported patch operations, see [Patch Operations](https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html).")
	apigateway_updateStageCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_updateStageCmd.Flags().String("stage-name", "", "The name of the Stage resource to change information about.")
	apigateway_updateStageCmd.MarkFlagRequired("rest-api-id")
	apigateway_updateStageCmd.MarkFlagRequired("stage-name")
	apigatewayCmd.AddCommand(apigateway_updateStageCmd)
}
