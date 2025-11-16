package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcs_createClusterCmd = &cobra.Command{
	Use:   "create-cluster",
	Short: "Creates a cluster in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcs_createClusterCmd).Standalone()

	pcs_createClusterCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	pcs_createClusterCmd.Flags().String("cluster-name", "", "A name to identify the cluster.")
	pcs_createClusterCmd.Flags().String("networking", "", "The networking configuration used to set up the cluster's control plane.")
	pcs_createClusterCmd.Flags().String("scheduler", "", "The cluster management and job scheduling software associated with the cluster.")
	pcs_createClusterCmd.Flags().String("size", "", "A value that determines the maximum number of compute nodes in the cluster and the maximum number of jobs (active and queued).")
	pcs_createClusterCmd.Flags().String("slurm-configuration", "", "Additional options related to the Slurm scheduler.")
	pcs_createClusterCmd.Flags().String("tags", "", "1 or more tags added to the resource.")
	pcs_createClusterCmd.MarkFlagRequired("cluster-name")
	pcs_createClusterCmd.MarkFlagRequired("networking")
	pcs_createClusterCmd.MarkFlagRequired("scheduler")
	pcs_createClusterCmd.MarkFlagRequired("size")
	pcsCmd.AddCommand(pcs_createClusterCmd)
}
