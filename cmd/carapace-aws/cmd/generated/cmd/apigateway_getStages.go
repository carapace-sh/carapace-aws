package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getStagesCmd = &cobra.Command{
	Use:   "get-stages",
	Short: "Gets information about one or more Stage resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getStagesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_getStagesCmd).Standalone()

		apigateway_getStagesCmd.Flags().String("deployment-id", "", "The stages' deployment identifiers.")
		apigateway_getStagesCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_getStagesCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_getStagesCmd)
}
