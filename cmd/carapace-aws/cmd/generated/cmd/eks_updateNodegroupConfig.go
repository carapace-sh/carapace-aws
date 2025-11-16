package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_updateNodegroupConfigCmd = &cobra.Command{
	Use:   "update-nodegroup-config",
	Short: "Updates an Amazon EKS managed node group configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_updateNodegroupConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_updateNodegroupConfigCmd).Standalone()

		eks_updateNodegroupConfigCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		eks_updateNodegroupConfigCmd.Flags().String("cluster-name", "", "The name of your cluster.")
		eks_updateNodegroupConfigCmd.Flags().String("labels", "", "The Kubernetes `labels` to apply to the nodes in the node group after the update.")
		eks_updateNodegroupConfigCmd.Flags().String("node-repair-config", "", "The node auto repair configuration for the node group.")
		eks_updateNodegroupConfigCmd.Flags().String("nodegroup-name", "", "The name of the managed node group to update.")
		eks_updateNodegroupConfigCmd.Flags().String("scaling-config", "", "The scaling configuration details for the Auto Scaling group after the update.")
		eks_updateNodegroupConfigCmd.Flags().String("taints", "", "The Kubernetes taints to be applied to the nodes in the node group after the update.")
		eks_updateNodegroupConfigCmd.Flags().String("update-config", "", "The node group update configuration.")
		eks_updateNodegroupConfigCmd.MarkFlagRequired("cluster-name")
		eks_updateNodegroupConfigCmd.MarkFlagRequired("nodegroup-name")
	})
	eksCmd.AddCommand(eks_updateNodegroupConfigCmd)
}
