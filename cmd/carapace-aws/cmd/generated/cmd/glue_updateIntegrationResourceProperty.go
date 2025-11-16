package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_updateIntegrationResourcePropertyCmd = &cobra.Command{
	Use:   "update-integration-resource-property",
	Short: "This API can be used for updating the `ResourceProperty` of the Glue connection (for the source) or Glue database ARN (for the target).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_updateIntegrationResourcePropertyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_updateIntegrationResourcePropertyCmd).Standalone()

		glue_updateIntegrationResourcePropertyCmd.Flags().String("resource-arn", "", "The connection ARN of the source, or the database ARN of the target.")
		glue_updateIntegrationResourcePropertyCmd.Flags().String("source-processing-properties", "", "The resource properties associated with the integration source.")
		glue_updateIntegrationResourcePropertyCmd.Flags().String("target-processing-properties", "", "The resource properties associated with the integration target.")
		glue_updateIntegrationResourcePropertyCmd.MarkFlagRequired("resource-arn")
	})
	glueCmd.AddCommand(glue_updateIntegrationResourcePropertyCmd)
}
