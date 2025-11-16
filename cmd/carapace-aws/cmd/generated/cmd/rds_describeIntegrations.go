package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeIntegrationsCmd = &cobra.Command{
	Use:   "describe-integrations",
	Short: "Describe one or more zero-ETL integrations with Amazon Redshift.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeIntegrationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_describeIntegrationsCmd).Standalone()

		rds_describeIntegrationsCmd.Flags().String("filters", "", "A filter that specifies one or more resources to return.")
		rds_describeIntegrationsCmd.Flags().String("integration-identifier", "", "The unique identifier of the integration.")
		rds_describeIntegrationsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeIntegrations` request.")
		rds_describeIntegrationsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	})
	rdsCmd.AddCommand(rds_describeIntegrationsCmd)
}
