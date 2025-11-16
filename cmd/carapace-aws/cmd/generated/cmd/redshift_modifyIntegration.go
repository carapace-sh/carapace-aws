package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_modifyIntegrationCmd = &cobra.Command{
	Use:   "modify-integration",
	Short: "Modifies a zero-ETL integration or S3 event integration with Amazon Redshift.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_modifyIntegrationCmd).Standalone()

	redshift_modifyIntegrationCmd.Flags().String("description", "", "A new description for the integration.")
	redshift_modifyIntegrationCmd.Flags().String("integration-arn", "", "The unique identifier of the integration to modify.")
	redshift_modifyIntegrationCmd.Flags().String("integration-name", "", "A new name for the integration.")
	redshift_modifyIntegrationCmd.MarkFlagRequired("integration-arn")
	redshiftCmd.AddCommand(redshift_modifyIntegrationCmd)
}
