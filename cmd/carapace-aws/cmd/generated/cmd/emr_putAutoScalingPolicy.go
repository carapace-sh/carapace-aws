package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_putAutoScalingPolicyCmd = &cobra.Command{
	Use:   "put-auto-scaling-policy",
	Short: "Creates or updates an automatic scaling policy for a core instance group or task instance group in an Amazon EMR cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_putAutoScalingPolicyCmd).Standalone()

	emr_putAutoScalingPolicyCmd.Flags().String("auto-scaling-policy", "", "Specifies the definition of the automatic scaling policy.")
	emr_putAutoScalingPolicyCmd.Flags().String("cluster-id", "", "Specifies the ID of a cluster.")
	emr_putAutoScalingPolicyCmd.Flags().String("instance-group-id", "", "Specifies the ID of the instance group to which the automatic scaling policy is applied.")
	emr_putAutoScalingPolicyCmd.MarkFlagRequired("auto-scaling-policy")
	emr_putAutoScalingPolicyCmd.MarkFlagRequired("cluster-id")
	emr_putAutoScalingPolicyCmd.MarkFlagRequired("instance-group-id")
	emrCmd.AddCommand(emr_putAutoScalingPolicyCmd)
}
