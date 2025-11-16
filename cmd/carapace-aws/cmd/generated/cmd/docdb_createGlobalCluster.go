package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_createGlobalClusterCmd = &cobra.Command{
	Use:   "create-global-cluster",
	Short: "Creates an Amazon DocumentDB global cluster that can span multiple multiple Amazon Web Services Regions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_createGlobalClusterCmd).Standalone()

	docdb_createGlobalClusterCmd.Flags().String("database-name", "", "The name for your database of up to 64 alpha-numeric characters.")
	docdb_createGlobalClusterCmd.Flags().String("deletion-protection", "", "The deletion protection setting for the new global cluster.")
	docdb_createGlobalClusterCmd.Flags().String("engine", "", "The name of the database engine to be used for this cluster.")
	docdb_createGlobalClusterCmd.Flags().String("engine-version", "", "The engine version of the global cluster.")
	docdb_createGlobalClusterCmd.Flags().String("global-cluster-identifier", "", "The cluster identifier of the new global cluster.")
	docdb_createGlobalClusterCmd.Flags().String("source-dbcluster-identifier", "", "The Amazon Resource Name (ARN) to use as the primary cluster of the global cluster.")
	docdb_createGlobalClusterCmd.Flags().String("storage-encrypted", "", "The storage encryption setting for the new global cluster.")
	docdb_createGlobalClusterCmd.MarkFlagRequired("global-cluster-identifier")
	docdbCmd.AddCommand(docdb_createGlobalClusterCmd)
}
