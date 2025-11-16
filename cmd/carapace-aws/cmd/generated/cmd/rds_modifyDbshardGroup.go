package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_modifyDbshardGroupCmd = &cobra.Command{
	Use:   "modify-dbshard-group",
	Short: "Modifies the settings of an Aurora Limitless Database DB shard group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_modifyDbshardGroupCmd).Standalone()

	rds_modifyDbshardGroupCmd.Flags().String("compute-redundancy", "", "Specifies whether to create standby DB shard groups for the DB shard group.")
	rds_modifyDbshardGroupCmd.Flags().String("dbshard-group-identifier", "", "The name of the DB shard group to modify.")
	rds_modifyDbshardGroupCmd.Flags().String("max-acu", "", "The maximum capacity of the DB shard group in Aurora capacity units (ACUs).")
	rds_modifyDbshardGroupCmd.Flags().String("min-acu", "", "The minimum capacity of the DB shard group in Aurora capacity units (ACUs).")
	rds_modifyDbshardGroupCmd.MarkFlagRequired("dbshard-group-identifier")
	rdsCmd.AddCommand(rds_modifyDbshardGroupCmd)
}
