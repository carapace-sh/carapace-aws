package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_deleteGlobalClusterCmd = &cobra.Command{
	Use:   "delete-global-cluster",
	Short: "Deletes a global database cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_deleteGlobalClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_deleteGlobalClusterCmd).Standalone()

		rds_deleteGlobalClusterCmd.Flags().String("global-cluster-identifier", "", "The cluster identifier of the global database cluster being deleted.")
		rds_deleteGlobalClusterCmd.MarkFlagRequired("global-cluster-identifier")
	})
	rdsCmd.AddCommand(rds_deleteGlobalClusterCmd)
}
