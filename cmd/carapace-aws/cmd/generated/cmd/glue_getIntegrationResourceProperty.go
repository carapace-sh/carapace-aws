package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getIntegrationResourcePropertyCmd = &cobra.Command{
	Use:   "get-integration-resource-property",
	Short: "This API is used for fetching the `ResourceProperty` of the Glue connection (for the source) or Glue database ARN (for the target)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getIntegrationResourcePropertyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getIntegrationResourcePropertyCmd).Standalone()

		glue_getIntegrationResourcePropertyCmd.Flags().String("resource-arn", "", "The connection ARN of the source, or the database ARN of the target.")
		glue_getIntegrationResourcePropertyCmd.MarkFlagRequired("resource-arn")
	})
	glueCmd.AddCommand(glue_getIntegrationResourcePropertyCmd)
}
