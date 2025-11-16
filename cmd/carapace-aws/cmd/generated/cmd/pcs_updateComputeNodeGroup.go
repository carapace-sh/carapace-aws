package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcs_updateComputeNodeGroupCmd = &cobra.Command{
	Use:   "update-compute-node-group",
	Short: "Updates a compute node group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcs_updateComputeNodeGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcs_updateComputeNodeGroupCmd).Standalone()

		pcs_updateComputeNodeGroupCmd.Flags().String("ami-id", "", "The ID of the Amazon Machine Image (AMI) that PCS uses to launch instances.")
		pcs_updateComputeNodeGroupCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		pcs_updateComputeNodeGroupCmd.Flags().String("cluster-identifier", "", "The name or ID of the cluster of the compute node group.")
		pcs_updateComputeNodeGroupCmd.Flags().String("compute-node-group-identifier", "", "The name or ID of the compute node group.")
		pcs_updateComputeNodeGroupCmd.Flags().String("custom-launch-template", "", "")
		pcs_updateComputeNodeGroupCmd.Flags().String("iam-instance-profile-arn", "", "The Amazon Resource Name (ARN) of the IAM instance profile used to pass an IAM role when launching EC2 instances.")
		pcs_updateComputeNodeGroupCmd.Flags().String("purchase-option", "", "Specifies how EC2 instances are purchased on your behalf.")
		pcs_updateComputeNodeGroupCmd.Flags().String("scaling-configuration", "", "Specifies the boundaries of the compute node group auto scaling.")
		pcs_updateComputeNodeGroupCmd.Flags().String("slurm-configuration", "", "Additional options related to the Slurm scheduler.")
		pcs_updateComputeNodeGroupCmd.Flags().String("spot-options", "", "")
		pcs_updateComputeNodeGroupCmd.Flags().String("subnet-ids", "", "The list of subnet IDs where the compute node group provisions instances.")
		pcs_updateComputeNodeGroupCmd.MarkFlagRequired("cluster-identifier")
		pcs_updateComputeNodeGroupCmd.MarkFlagRequired("compute-node-group-identifier")
	})
	pcsCmd.AddCommand(pcs_updateComputeNodeGroupCmd)
}
