package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_rebootDbshardGroupCmd = &cobra.Command{
	Use:   "reboot-dbshard-group",
	Short: "You might need to reboot your DB shard group, usually for maintenance reasons.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_rebootDbshardGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_rebootDbshardGroupCmd).Standalone()

		rds_rebootDbshardGroupCmd.Flags().String("dbshard-group-identifier", "", "The name of the DB shard group to reboot.")
		rds_rebootDbshardGroupCmd.MarkFlagRequired("dbshard-group-identifier")
	})
	rdsCmd.AddCommand(rds_rebootDbshardGroupCmd)
}
