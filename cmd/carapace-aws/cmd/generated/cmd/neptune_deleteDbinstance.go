package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_deleteDbinstanceCmd = &cobra.Command{
	Use:   "delete-dbinstance",
	Short: "The DeleteDBInstance action deletes a previously provisioned DB instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_deleteDbinstanceCmd).Standalone()

	neptune_deleteDbinstanceCmd.Flags().String("dbinstance-identifier", "", "The DB instance identifier for the DB instance to be deleted.")
	neptune_deleteDbinstanceCmd.Flags().String("final-dbsnapshot-identifier", "", "The DBSnapshotIdentifier of the new DBSnapshot created when SkipFinalSnapshot is set to `false`.")
	neptune_deleteDbinstanceCmd.Flags().Bool("no-skip-final-snapshot", false, "Determines whether a final DB snapshot is created before the DB instance is deleted.")
	neptune_deleteDbinstanceCmd.Flags().Bool("skip-final-snapshot", false, "Determines whether a final DB snapshot is created before the DB instance is deleted.")
	neptune_deleteDbinstanceCmd.MarkFlagRequired("dbinstance-identifier")
	neptune_deleteDbinstanceCmd.Flag("no-skip-final-snapshot").Hidden = true
	neptuneCmd.AddCommand(neptune_deleteDbinstanceCmd)
}
