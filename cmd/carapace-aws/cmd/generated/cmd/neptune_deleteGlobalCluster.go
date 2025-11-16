package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_deleteGlobalClusterCmd = &cobra.Command{
	Use:   "delete-global-cluster",
	Short: "Deletes a global database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_deleteGlobalClusterCmd).Standalone()

	neptune_deleteGlobalClusterCmd.Flags().String("global-cluster-identifier", "", "The cluster identifier of the global database cluster being deleted.")
	neptune_deleteGlobalClusterCmd.MarkFlagRequired("global-cluster-identifier")
	neptuneCmd.AddCommand(neptune_deleteGlobalClusterCmd)
}
