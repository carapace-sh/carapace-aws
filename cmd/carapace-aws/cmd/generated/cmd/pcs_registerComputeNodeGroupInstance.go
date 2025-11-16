package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcs_registerComputeNodeGroupInstanceCmd = &cobra.Command{
	Use:   "register-compute-node-group-instance",
	Short: "This API action isn't intended for you to use.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcs_registerComputeNodeGroupInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcs_registerComputeNodeGroupInstanceCmd).Standalone()

		pcs_registerComputeNodeGroupInstanceCmd.Flags().String("bootstrap-id", "", "The client-generated token to allow for retries.")
		pcs_registerComputeNodeGroupInstanceCmd.Flags().String("cluster-identifier", "", "The name or ID of the cluster to register the compute node group instance in.")
		pcs_registerComputeNodeGroupInstanceCmd.MarkFlagRequired("bootstrap-id")
		pcs_registerComputeNodeGroupInstanceCmd.MarkFlagRequired("cluster-identifier")
	})
	pcsCmd.AddCommand(pcs_registerComputeNodeGroupInstanceCmd)
}
