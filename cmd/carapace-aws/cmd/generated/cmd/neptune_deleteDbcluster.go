package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_deleteDbclusterCmd = &cobra.Command{
	Use:   "delete-dbcluster",
	Short: "The DeleteDBCluster action deletes a previously provisioned DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_deleteDbclusterCmd).Standalone()

	neptune_deleteDbclusterCmd.Flags().String("dbcluster-identifier", "", "The DB cluster identifier for the DB cluster to be deleted.")
	neptune_deleteDbclusterCmd.Flags().String("final-dbsnapshot-identifier", "", "The DB cluster snapshot identifier of the new DB cluster snapshot created when `SkipFinalSnapshot` is set to `false`.")
	neptune_deleteDbclusterCmd.Flags().Bool("no-skip-final-snapshot", false, "Determines whether a final DB cluster snapshot is created before the DB cluster is deleted.")
	neptune_deleteDbclusterCmd.Flags().Bool("skip-final-snapshot", false, "Determines whether a final DB cluster snapshot is created before the DB cluster is deleted.")
	neptune_deleteDbclusterCmd.MarkFlagRequired("dbcluster-identifier")
	neptune_deleteDbclusterCmd.Flag("no-skip-final-snapshot").Hidden = true
	neptuneCmd.AddCommand(neptune_deleteDbclusterCmd)
}
