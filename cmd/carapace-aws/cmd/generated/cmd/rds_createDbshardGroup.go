package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_createDbshardGroupCmd = &cobra.Command{
	Use:   "create-dbshard-group",
	Short: "Creates a new DB shard group for Aurora Limitless Database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_createDbshardGroupCmd).Standalone()

	rds_createDbshardGroupCmd.Flags().String("compute-redundancy", "", "Specifies whether to create standby standby DB data access shard for the DB shard group.")
	rds_createDbshardGroupCmd.Flags().String("dbcluster-identifier", "", "The name of the primary DB cluster for the DB shard group.")
	rds_createDbshardGroupCmd.Flags().String("dbshard-group-identifier", "", "The name of the DB shard group.")
	rds_createDbshardGroupCmd.Flags().String("max-acu", "", "The maximum capacity of the DB shard group in Aurora capacity units (ACUs).")
	rds_createDbshardGroupCmd.Flags().String("min-acu", "", "The minimum capacity of the DB shard group in Aurora capacity units (ACUs).")
	rds_createDbshardGroupCmd.Flags().String("publicly-accessible", "", "Specifies whether the DB shard group is publicly accessible.")
	rds_createDbshardGroupCmd.Flags().String("tags", "", "")
	rds_createDbshardGroupCmd.MarkFlagRequired("dbcluster-identifier")
	rds_createDbshardGroupCmd.MarkFlagRequired("dbshard-group-identifier")
	rds_createDbshardGroupCmd.MarkFlagRequired("max-acu")
	rdsCmd.AddCommand(rds_createDbshardGroupCmd)
}
