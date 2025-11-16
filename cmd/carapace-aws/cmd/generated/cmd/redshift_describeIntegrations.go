package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeIntegrationsCmd = &cobra.Command{
	Use:   "describe-integrations",
	Short: "Describes one or more zero-ETL or S3 event integrations with Amazon Redshift.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeIntegrationsCmd).Standalone()

	redshift_describeIntegrationsCmd.Flags().String("filters", "", "A filter that specifies one or more resources to return.")
	redshift_describeIntegrationsCmd.Flags().String("integration-arn", "", "The unique identifier of the integration.")
	redshift_describeIntegrationsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeIntegrations` request.")
	redshift_describeIntegrationsCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
	redshiftCmd.AddCommand(redshift_describeIntegrationsCmd)
}
