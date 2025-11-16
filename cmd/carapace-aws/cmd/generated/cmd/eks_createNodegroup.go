package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_createNodegroupCmd = &cobra.Command{
	Use:   "create-nodegroup",
	Short: "Creates a managed node group for an Amazon EKS cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_createNodegroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_createNodegroupCmd).Standalone()

		eks_createNodegroupCmd.Flags().String("ami-type", "", "The AMI type for your node group.")
		eks_createNodegroupCmd.Flags().String("capacity-type", "", "The capacity type for your node group.")
		eks_createNodegroupCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		eks_createNodegroupCmd.Flags().String("cluster-name", "", "The name of your cluster.")
		eks_createNodegroupCmd.Flags().String("disk-size", "", "The root device disk size (in GiB) for your node group instances.")
		eks_createNodegroupCmd.Flags().String("instance-types", "", "Specify the instance types for a node group.")
		eks_createNodegroupCmd.Flags().String("labels", "", "The Kubernetes `labels` to apply to the nodes in the node group when they are created.")
		eks_createNodegroupCmd.Flags().String("launch-template", "", "An object representing a node group's launch template specification.")
		eks_createNodegroupCmd.Flags().String("node-repair-config", "", "The node auto repair configuration for the node group.")
		eks_createNodegroupCmd.Flags().String("node-role", "", "The Amazon Resource Name (ARN) of the IAM role to associate with your node group.")
		eks_createNodegroupCmd.Flags().String("nodegroup-name", "", "The unique name to give your node group.")
		eks_createNodegroupCmd.Flags().String("release-version", "", "The AMI version of the Amazon EKS optimized AMI to use with your node group.")
		eks_createNodegroupCmd.Flags().String("remote-access", "", "The remote access configuration to use with your node group.")
		eks_createNodegroupCmd.Flags().String("scaling-config", "", "The scaling configuration details for the Auto Scaling group that is created for your node group.")
		eks_createNodegroupCmd.Flags().String("subnets", "", "The subnets to use for the Auto Scaling group that is created for your node group.")
		eks_createNodegroupCmd.Flags().String("tags", "", "Metadata that assists with categorization and organization.")
		eks_createNodegroupCmd.Flags().String("taints", "", "The Kubernetes taints to be applied to the nodes in the node group.")
		eks_createNodegroupCmd.Flags().String("update-config", "", "The node group update configuration.")
		eks_createNodegroupCmd.Flags().String("version", "", "The Kubernetes version to use for your managed nodes.")
		eks_createNodegroupCmd.MarkFlagRequired("cluster-name")
		eks_createNodegroupCmd.MarkFlagRequired("node-role")
		eks_createNodegroupCmd.MarkFlagRequired("nodegroup-name")
		eks_createNodegroupCmd.MarkFlagRequired("subnets")
	})
	eksCmd.AddCommand(eks_createNodegroupCmd)
}
