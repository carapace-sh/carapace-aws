package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_modifyIntegrationCmd = &cobra.Command{
	Use:   "modify-integration",
	Short: "Modifies a zero-ETL integration with Amazon Redshift.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_modifyIntegrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_modifyIntegrationCmd).Standalone()

		rds_modifyIntegrationCmd.Flags().String("data-filter", "", "A new data filter for the integration.")
		rds_modifyIntegrationCmd.Flags().String("description", "", "A new description for the integration.")
		rds_modifyIntegrationCmd.Flags().String("integration-identifier", "", "The unique identifier of the integration to modify.")
		rds_modifyIntegrationCmd.Flags().String("integration-name", "", "A new name for the integration.")
		rds_modifyIntegrationCmd.MarkFlagRequired("integration-identifier")
	})
	rdsCmd.AddCommand(rds_modifyIntegrationCmd)
}
