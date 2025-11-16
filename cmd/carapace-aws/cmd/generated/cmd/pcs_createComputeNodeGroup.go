package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcs_createComputeNodeGroupCmd = &cobra.Command{
	Use:   "create-compute-node-group",
	Short: "Creates a managed set of compute nodes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcs_createComputeNodeGroupCmd).Standalone()

	pcs_createComputeNodeGroupCmd.Flags().String("ami-id", "", "The ID of the Amazon Machine Image (AMI) that PCS uses to launch compute nodes (Amazon EC2 instances).")
	pcs_createComputeNodeGroupCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	pcs_createComputeNodeGroupCmd.Flags().String("cluster-identifier", "", "The name or ID of the cluster to create a compute node group in.")
	pcs_createComputeNodeGroupCmd.Flags().String("compute-node-group-name", "", "A name to identify the cluster.")
	pcs_createComputeNodeGroupCmd.Flags().String("custom-launch-template", "", "")
	pcs_createComputeNodeGroupCmd.Flags().String("iam-instance-profile-arn", "", "The Amazon Resource Name (ARN) of the IAM instance profile used to pass an IAM role when launching EC2 instances.")
	pcs_createComputeNodeGroupCmd.Flags().String("instance-configs", "", "A list of EC2 instance configurations that PCS can provision in the compute node group.")
	pcs_createComputeNodeGroupCmd.Flags().String("purchase-option", "", "Specifies how EC2 instances are purchased on your behalf.")
	pcs_createComputeNodeGroupCmd.Flags().String("scaling-configuration", "", "Specifies the boundaries of the compute node group auto scaling.")
	pcs_createComputeNodeGroupCmd.Flags().String("slurm-configuration", "", "Additional options related to the Slurm scheduler.")
	pcs_createComputeNodeGroupCmd.Flags().String("spot-options", "", "")
	pcs_createComputeNodeGroupCmd.Flags().String("subnet-ids", "", "The list of subnet IDs where the compute node group launches instances.")
	pcs_createComputeNodeGroupCmd.Flags().String("tags", "", "1 or more tags added to the resource.")
	pcs_createComputeNodeGroupCmd.MarkFlagRequired("cluster-identifier")
	pcs_createComputeNodeGroupCmd.MarkFlagRequired("compute-node-group-name")
	pcs_createComputeNodeGroupCmd.MarkFlagRequired("custom-launch-template")
	pcs_createComputeNodeGroupCmd.MarkFlagRequired("iam-instance-profile-arn")
	pcs_createComputeNodeGroupCmd.MarkFlagRequired("instance-configs")
	pcs_createComputeNodeGroupCmd.MarkFlagRequired("scaling-configuration")
	pcs_createComputeNodeGroupCmd.MarkFlagRequired("subnet-ids")
	pcsCmd.AddCommand(pcs_createComputeNodeGroupCmd)
}
