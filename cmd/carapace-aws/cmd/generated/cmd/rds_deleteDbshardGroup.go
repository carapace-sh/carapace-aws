package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_deleteDbshardGroupCmd = &cobra.Command{
	Use:   "delete-dbshard-group",
	Short: "Deletes an Aurora Limitless Database DB shard group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_deleteDbshardGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_deleteDbshardGroupCmd).Standalone()

		rds_deleteDbshardGroupCmd.Flags().String("dbshard-group-identifier", "", "The name of the DB shard group to delete.")
		rds_deleteDbshardGroupCmd.MarkFlagRequired("dbshard-group-identifier")
	})
	rdsCmd.AddCommand(rds_deleteDbshardGroupCmd)
}
