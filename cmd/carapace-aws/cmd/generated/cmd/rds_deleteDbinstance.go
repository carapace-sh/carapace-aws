package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_deleteDbinstanceCmd = &cobra.Command{
	Use:   "delete-dbinstance",
	Short: "Deletes a previously provisioned DB instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_deleteDbinstanceCmd).Standalone()

	rds_deleteDbinstanceCmd.Flags().String("dbinstance-identifier", "", "The DB instance identifier for the DB instance to be deleted.")
	rds_deleteDbinstanceCmd.Flags().String("delete-automated-backups", "", "Specifies whether to remove automated backups immediately after the DB instance is deleted.")
	rds_deleteDbinstanceCmd.Flags().String("final-dbsnapshot-identifier", "", "The `DBSnapshotIdentifier` of the new `DBSnapshot` created when the `SkipFinalSnapshot` parameter is disabled.")
	rds_deleteDbinstanceCmd.Flags().Bool("no-skip-final-snapshot", false, "Specifies whether to skip the creation of a final DB snapshot before deleting the instance.")
	rds_deleteDbinstanceCmd.Flags().Bool("skip-final-snapshot", false, "Specifies whether to skip the creation of a final DB snapshot before deleting the instance.")
	rds_deleteDbinstanceCmd.MarkFlagRequired("dbinstance-identifier")
	rds_deleteDbinstanceCmd.Flag("no-skip-final-snapshot").Hidden = true
	rdsCmd.AddCommand(rds_deleteDbinstanceCmd)
}
