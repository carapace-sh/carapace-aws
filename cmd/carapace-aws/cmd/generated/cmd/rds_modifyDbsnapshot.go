package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_modifyDbsnapshotCmd = &cobra.Command{
	Use:   "modify-dbsnapshot",
	Short: "Updates a manual DB snapshot with a new engine version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_modifyDbsnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_modifyDbsnapshotCmd).Standalone()

		rds_modifyDbsnapshotCmd.Flags().String("dbsnapshot-identifier", "", "The identifier of the DB snapshot to modify.")
		rds_modifyDbsnapshotCmd.Flags().String("engine-version", "", "The engine version to upgrade the DB snapshot to.")
		rds_modifyDbsnapshotCmd.Flags().String("option-group-name", "", "The option group to identify with the upgraded DB snapshot.")
		rds_modifyDbsnapshotCmd.MarkFlagRequired("dbsnapshot-identifier")
	})
	rdsCmd.AddCommand(rds_modifyDbsnapshotCmd)
}
