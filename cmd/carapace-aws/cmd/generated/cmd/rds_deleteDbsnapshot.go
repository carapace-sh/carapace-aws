package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_deleteDbsnapshotCmd = &cobra.Command{
	Use:   "delete-dbsnapshot",
	Short: "Deletes a DB snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_deleteDbsnapshotCmd).Standalone()

	rds_deleteDbsnapshotCmd.Flags().String("dbsnapshot-identifier", "", "The DB snapshot identifier.")
	rds_deleteDbsnapshotCmd.MarkFlagRequired("dbsnapshot-identifier")
	rdsCmd.AddCommand(rds_deleteDbsnapshotCmd)
}
