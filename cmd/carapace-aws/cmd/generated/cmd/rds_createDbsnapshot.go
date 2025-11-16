package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_createDbsnapshotCmd = &cobra.Command{
	Use:   "create-dbsnapshot",
	Short: "Creates a snapshot of a DB instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_createDbsnapshotCmd).Standalone()

	rds_createDbsnapshotCmd.Flags().String("dbinstance-identifier", "", "The identifier of the DB instance that you want to create the snapshot of.")
	rds_createDbsnapshotCmd.Flags().String("dbsnapshot-identifier", "", "The identifier for the DB snapshot.")
	rds_createDbsnapshotCmd.Flags().String("tags", "", "")
	rds_createDbsnapshotCmd.MarkFlagRequired("dbinstance-identifier")
	rds_createDbsnapshotCmd.MarkFlagRequired("dbsnapshot-identifier")
	rdsCmd.AddCommand(rds_createDbsnapshotCmd)
}
