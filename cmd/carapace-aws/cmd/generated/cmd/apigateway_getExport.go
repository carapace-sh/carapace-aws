package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getExportCmd = &cobra.Command{
	Use:   "get-export",
	Short: "Exports a deployed version of a RestApi in a specified format.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getExportCmd).Standalone()

	apigateway_getExportCmd.Flags().String("accepts", "", "The content-type of the export, for example `application/json`.")
	apigateway_getExportCmd.Flags().String("export-type", "", "The type of export.")
	apigateway_getExportCmd.Flags().String("parameters", "", "A key-value map of query string parameters that specify properties of the export, depending on the requested `exportType`.")
	apigateway_getExportCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_getExportCmd.Flags().String("stage-name", "", "The name of the Stage that will be exported.")
	apigateway_getExportCmd.MarkFlagRequired("export-type")
	apigateway_getExportCmd.MarkFlagRequired("rest-api-id")
	apigateway_getExportCmd.MarkFlagRequired("stage-name")
	apigatewayCmd.AddCommand(apigateway_getExportCmd)
}
