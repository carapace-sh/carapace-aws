package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_deleteIntegrationCmd = &cobra.Command{
	Use:   "delete-integration",
	Short: "Deletes a zero-ETL integration with Amazon Redshift.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_deleteIntegrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_deleteIntegrationCmd).Standalone()

		rds_deleteIntegrationCmd.Flags().String("integration-identifier", "", "The unique identifier of the integration.")
		rds_deleteIntegrationCmd.MarkFlagRequired("integration-identifier")
	})
	rdsCmd.AddCommand(rds_deleteIntegrationCmd)
}
