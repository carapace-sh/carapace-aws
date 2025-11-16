package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_removeAutoScalingPolicyCmd = &cobra.Command{
	Use:   "remove-auto-scaling-policy",
	Short: "Removes an automatic scaling policy from a specified instance group within an Amazon EMR cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_removeAutoScalingPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_removeAutoScalingPolicyCmd).Standalone()

		emr_removeAutoScalingPolicyCmd.Flags().String("cluster-id", "", "Specifies the ID of a cluster.")
		emr_removeAutoScalingPolicyCmd.Flags().String("instance-group-id", "", "Specifies the ID of the instance group to which the scaling policy is applied.")
		emr_removeAutoScalingPolicyCmd.MarkFlagRequired("cluster-id")
		emr_removeAutoScalingPolicyCmd.MarkFlagRequired("instance-group-id")
	})
	emrCmd.AddCommand(emr_removeAutoScalingPolicyCmd)
}
