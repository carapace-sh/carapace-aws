package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_modifyGlobalClusterCmd = &cobra.Command{
	Use:   "modify-global-cluster",
	Short: "Modify a setting for an Amazon DocumentDB global cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_modifyGlobalClusterCmd).Standalone()

	docdb_modifyGlobalClusterCmd.Flags().String("deletion-protection", "", "Indicates if the global cluster has deletion protection enabled.")
	docdb_modifyGlobalClusterCmd.Flags().String("global-cluster-identifier", "", "The identifier for the global cluster being modified.")
	docdb_modifyGlobalClusterCmd.Flags().String("new-global-cluster-identifier", "", "The new identifier for a global cluster when you modify a global cluster.")
	docdb_modifyGlobalClusterCmd.MarkFlagRequired("global-cluster-identifier")
	docdbCmd.AddCommand(docdb_modifyGlobalClusterCmd)
}
