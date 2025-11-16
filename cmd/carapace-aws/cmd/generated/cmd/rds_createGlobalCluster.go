package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_createGlobalClusterCmd = &cobra.Command{
	Use:   "create-global-cluster",
	Short: "Creates an Aurora global database spread across multiple Amazon Web Services Regions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_createGlobalClusterCmd).Standalone()

	rds_createGlobalClusterCmd.Flags().String("database-name", "", "The name for your database of up to 64 alphanumeric characters.")
	rds_createGlobalClusterCmd.Flags().String("deletion-protection", "", "Specifies whether to enable deletion protection for the new global database cluster.")
	rds_createGlobalClusterCmd.Flags().String("engine", "", "The database engine to use for this global database cluster.")
	rds_createGlobalClusterCmd.Flags().String("engine-lifecycle-support", "", "The life cycle type for this global database cluster.")
	rds_createGlobalClusterCmd.Flags().String("engine-version", "", "The engine version to use for this global database cluster.")
	rds_createGlobalClusterCmd.Flags().String("global-cluster-identifier", "", "The cluster identifier for this global database cluster.")
	rds_createGlobalClusterCmd.Flags().String("source-dbcluster-identifier", "", "The Amazon Resource Name (ARN) to use as the primary cluster of the global database.")
	rds_createGlobalClusterCmd.Flags().String("storage-encrypted", "", "Specifies whether to enable storage encryption for the new global database cluster.")
	rds_createGlobalClusterCmd.Flags().String("tags", "", "Tags to assign to the global cluster.")
	rds_createGlobalClusterCmd.MarkFlagRequired("global-cluster-identifier")
	rdsCmd.AddCommand(rds_createGlobalClusterCmd)
}
