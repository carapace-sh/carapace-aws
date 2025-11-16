package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_modifyIntegrationCmd = &cobra.Command{
	Use:   "modify-integration",
	Short: "Modifies a Zero-ETL integration in the caller's account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_modifyIntegrationCmd).Standalone()

	glue_modifyIntegrationCmd.Flags().String("data-filter", "", "Selects source tables for the integration using Maxwell filter syntax.")
	glue_modifyIntegrationCmd.Flags().String("description", "", "A description of the integration.")
	glue_modifyIntegrationCmd.Flags().String("integration-config", "", "")
	glue_modifyIntegrationCmd.Flags().String("integration-identifier", "", "The Amazon Resource Name (ARN) for the integration.")
	glue_modifyIntegrationCmd.Flags().String("integration-name", "", "A unique name for an integration in Glue.")
	glue_modifyIntegrationCmd.MarkFlagRequired("integration-identifier")
	glueCmd.AddCommand(glue_modifyIntegrationCmd)
}
