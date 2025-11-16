package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_deleteDbclusterCmd = &cobra.Command{
	Use:   "delete-dbcluster",
	Short: "The DeleteDBCluster action deletes a previously provisioned DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_deleteDbclusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_deleteDbclusterCmd).Standalone()

		rds_deleteDbclusterCmd.Flags().String("dbcluster-identifier", "", "The DB cluster identifier for the DB cluster to be deleted.")
		rds_deleteDbclusterCmd.Flags().String("delete-automated-backups", "", "Specifies whether to remove automated backups immediately after the DB cluster is deleted.")
		rds_deleteDbclusterCmd.Flags().String("final-dbsnapshot-identifier", "", "The DB cluster snapshot identifier of the new DB cluster snapshot created when `SkipFinalSnapshot` is disabled.")
		rds_deleteDbclusterCmd.Flags().Bool("no-skip-final-snapshot", false, "Specifies whether to skip the creation of a final DB cluster snapshot before RDS deletes the DB cluster.")
		rds_deleteDbclusterCmd.Flags().Bool("skip-final-snapshot", false, "Specifies whether to skip the creation of a final DB cluster snapshot before RDS deletes the DB cluster.")
		rds_deleteDbclusterCmd.MarkFlagRequired("dbcluster-identifier")
		rds_deleteDbclusterCmd.Flag("no-skip-final-snapshot").Hidden = true
	})
	rdsCmd.AddCommand(rds_deleteDbclusterCmd)
}
