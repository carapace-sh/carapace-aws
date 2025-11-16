package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeTableRestoreStatusCmd = &cobra.Command{
	Use:   "describe-table-restore-status",
	Short: "Lists the status of one or more table restore requests made using the [RestoreTableFromClusterSnapshot]() API action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeTableRestoreStatusCmd).Standalone()

	redshift_describeTableRestoreStatusCmd.Flags().String("cluster-identifier", "", "The Amazon Redshift cluster that the table is being restored to.")
	redshift_describeTableRestoreStatusCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeTableRestoreStatus` request.")
	redshift_describeTableRestoreStatusCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	redshift_describeTableRestoreStatusCmd.Flags().String("table-restore-request-id", "", "The identifier of the table restore request to return status for.")
	redshiftCmd.AddCommand(redshift_describeTableRestoreStatusCmd)
}
