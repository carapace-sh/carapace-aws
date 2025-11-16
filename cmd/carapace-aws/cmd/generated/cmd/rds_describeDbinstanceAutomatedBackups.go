package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeDbinstanceAutomatedBackupsCmd = &cobra.Command{
	Use:   "describe-dbinstance-automated-backups",
	Short: "Displays backups for both current and deleted instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeDbinstanceAutomatedBackupsCmd).Standalone()

	rds_describeDbinstanceAutomatedBackupsCmd.Flags().String("dbi-resource-id", "", "The resource ID of the DB instance that is the source of the automated backup.")
	rds_describeDbinstanceAutomatedBackupsCmd.Flags().String("dbinstance-automated-backups-arn", "", "The Amazon Resource Name (ARN) of the replicated automated backups, for example, `arn:aws:rds:us-east-1:123456789012:auto-backup:ab-L2IJCEXJP7XQ7HOJ4SIEXAMPLE`.")
	rds_describeDbinstanceAutomatedBackupsCmd.Flags().String("dbinstance-identifier", "", "(Optional) The user-supplied instance identifier.")
	rds_describeDbinstanceAutomatedBackupsCmd.Flags().String("filters", "", "A filter that specifies which resources to return based on status.")
	rds_describeDbinstanceAutomatedBackupsCmd.Flags().String("marker", "", "The pagination token provided in the previous request.")
	rds_describeDbinstanceAutomatedBackupsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	rdsCmd.AddCommand(rds_describeDbinstanceAutomatedBackupsCmd)
}
