package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_updateNodegroupVersionCmd = &cobra.Command{
	Use:   "update-nodegroup-version",
	Short: "Updates the Kubernetes version or AMI version of an Amazon EKS managed node group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_updateNodegroupVersionCmd).Standalone()

	eks_updateNodegroupVersionCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	eks_updateNodegroupVersionCmd.Flags().String("cluster-name", "", "The name of your cluster.")
	eks_updateNodegroupVersionCmd.Flags().Bool("force", false, "Force the update if any `Pod` on the existing node group can't be drained due to a `Pod` disruption budget issue.")
	eks_updateNodegroupVersionCmd.Flags().String("launch-template", "", "An object representing a node group's launch template specification.")
	eks_updateNodegroupVersionCmd.Flags().Bool("no-force", false, "Force the update if any `Pod` on the existing node group can't be drained due to a `Pod` disruption budget issue.")
	eks_updateNodegroupVersionCmd.Flags().String("nodegroup-name", "", "The name of the managed node group to update.")
	eks_updateNodegroupVersionCmd.Flags().String("release-version", "", "The AMI version of the Amazon EKS optimized AMI to use for the update.")
	eks_updateNodegroupVersionCmd.Flags().String("version", "", "The Kubernetes version to update to.")
	eks_updateNodegroupVersionCmd.MarkFlagRequired("cluster-name")
	eks_updateNodegroupVersionCmd.Flag("no-force").Hidden = true
	eks_updateNodegroupVersionCmd.MarkFlagRequired("nodegroup-name")
	eksCmd.AddCommand(eks_updateNodegroupVersionCmd)
}
