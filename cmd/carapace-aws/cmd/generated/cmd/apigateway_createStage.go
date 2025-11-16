package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_createStageCmd = &cobra.Command{
	Use:   "create-stage",
	Short: "Creates a new Stage resource that references a pre-existing Deployment for the API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_createStageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_createStageCmd).Standalone()

		apigateway_createStageCmd.Flags().Bool("cache-cluster-enabled", false, "Whether cache clustering is enabled for the stage.")
		apigateway_createStageCmd.Flags().String("cache-cluster-size", "", "The stage's cache capacity in GB.")
		apigateway_createStageCmd.Flags().String("canary-settings", "", "The canary deployment settings of this stage.")
		apigateway_createStageCmd.Flags().String("deployment-id", "", "The identifier of the Deployment resource for the Stage resource.")
		apigateway_createStageCmd.Flags().String("description", "", "The description of the Stage resource.")
		apigateway_createStageCmd.Flags().String("documentation-version", "", "The version of the associated API documentation.")
		apigateway_createStageCmd.Flags().Bool("no-cache-cluster-enabled", false, "Whether cache clustering is enabled for the stage.")
		apigateway_createStageCmd.Flags().Bool("no-tracing-enabled", false, "Specifies whether active tracing with X-ray is enabled for the Stage.")
		apigateway_createStageCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_createStageCmd.Flags().String("stage-name", "", "The name for the Stage resource.")
		apigateway_createStageCmd.Flags().String("tags", "", "The key-value map of strings.")
		apigateway_createStageCmd.Flags().Bool("tracing-enabled", false, "Specifies whether active tracing with X-ray is enabled for the Stage.")
		apigateway_createStageCmd.Flags().String("variables", "", "A map that defines the stage variables for the new Stage resource.")
		apigateway_createStageCmd.MarkFlagRequired("deployment-id")
		apigateway_createStageCmd.Flag("no-cache-cluster-enabled").Hidden = true
		apigateway_createStageCmd.Flag("no-tracing-enabled").Hidden = true
		apigateway_createStageCmd.MarkFlagRequired("rest-api-id")
		apigateway_createStageCmd.MarkFlagRequired("stage-name")
	})
	apigatewayCmd.AddCommand(apigateway_createStageCmd)
}
