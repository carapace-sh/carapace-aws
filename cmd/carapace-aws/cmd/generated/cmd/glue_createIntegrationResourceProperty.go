package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createIntegrationResourcePropertyCmd = &cobra.Command{
	Use:   "create-integration-resource-property",
	Short: "This API can be used for setting up the `ResourceProperty` of the Glue connection (for the source) or Glue database ARN (for the target).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createIntegrationResourcePropertyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_createIntegrationResourcePropertyCmd).Standalone()

		glue_createIntegrationResourcePropertyCmd.Flags().String("resource-arn", "", "The connection ARN of the source, or the database ARN of the target.")
		glue_createIntegrationResourcePropertyCmd.Flags().String("source-processing-properties", "", "The resource properties associated with the integration source.")
		glue_createIntegrationResourcePropertyCmd.Flags().String("target-processing-properties", "", "The resource properties associated with the integration target.")
		glue_createIntegrationResourcePropertyCmd.MarkFlagRequired("resource-arn")
	})
	glueCmd.AddCommand(glue_createIntegrationResourcePropertyCmd)
}
