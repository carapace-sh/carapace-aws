package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_deleteIntegrationCmd = &cobra.Command{
	Use:   "delete-integration",
	Short: "Deletes a zero-ETL integration or S3 event integration with Amazon Redshift.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_deleteIntegrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_deleteIntegrationCmd).Standalone()

		redshift_deleteIntegrationCmd.Flags().String("integration-arn", "", "The unique identifier of the integration to delete.")
		redshift_deleteIntegrationCmd.MarkFlagRequired("integration-arn")
	})
	redshiftCmd.AddCommand(redshift_deleteIntegrationCmd)
}
