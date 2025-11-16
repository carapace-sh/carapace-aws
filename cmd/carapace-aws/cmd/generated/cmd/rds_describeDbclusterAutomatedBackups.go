package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeDbclusterAutomatedBackupsCmd = &cobra.Command{
	Use:   "describe-dbcluster-automated-backups",
	Short: "Displays backups for both current and deleted DB clusters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeDbclusterAutomatedBackupsCmd).Standalone()

	rds_describeDbclusterAutomatedBackupsCmd.Flags().String("db-cluster-resource-id", "", "The resource ID of the DB cluster that is the source of the automated backup.")
	rds_describeDbclusterAutomatedBackupsCmd.Flags().String("dbcluster-identifier", "", "(Optional) The user-supplied DB cluster identifier.")
	rds_describeDbclusterAutomatedBackupsCmd.Flags().String("filters", "", "A filter that specifies which resources to return based on status.")
	rds_describeDbclusterAutomatedBackupsCmd.Flags().String("marker", "", "The pagination token provided in the previous request.")
	rds_describeDbclusterAutomatedBackupsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	rdsCmd.AddCommand(rds_describeDbclusterAutomatedBackupsCmd)
}
