package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcs_updateClusterCmd = &cobra.Command{
	Use:   "update-cluster",
	Short: "Updates a cluster configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcs_updateClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcs_updateClusterCmd).Standalone()

		pcs_updateClusterCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		pcs_updateClusterCmd.Flags().String("cluster-identifier", "", "The name or ID of the cluster to update.")
		pcs_updateClusterCmd.Flags().String("slurm-configuration", "", "Additional options related to the Slurm scheduler.")
		pcs_updateClusterCmd.MarkFlagRequired("cluster-identifier")
	})
	pcsCmd.AddCommand(pcs_updateClusterCmd)
}
