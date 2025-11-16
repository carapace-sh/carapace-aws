package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteIntegrationCmd = &cobra.Command{
	Use:   "delete-integration",
	Short: "Deletes the specified Zero-ETL integration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteIntegrationCmd).Standalone()

	glue_deleteIntegrationCmd.Flags().String("integration-identifier", "", "The Amazon Resource Name (ARN) for the integration.")
	glue_deleteIntegrationCmd.MarkFlagRequired("integration-identifier")
	glueCmd.AddCommand(glue_deleteIntegrationCmd)
}
