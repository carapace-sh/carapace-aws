package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_putManagedScalingPolicyCmd = &cobra.Command{
	Use:   "put-managed-scaling-policy",
	Short: "Creates or updates a managed scaling policy for an Amazon EMR cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_putManagedScalingPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_putManagedScalingPolicyCmd).Standalone()

		emr_putManagedScalingPolicyCmd.Flags().String("cluster-id", "", "Specifies the ID of an Amazon EMR cluster where the managed scaling policy is attached.")
		emr_putManagedScalingPolicyCmd.Flags().String("managed-scaling-policy", "", "Specifies the constraints for the managed scaling policy.")
		emr_putManagedScalingPolicyCmd.MarkFlagRequired("cluster-id")
		emr_putManagedScalingPolicyCmd.MarkFlagRequired("managed-scaling-policy")
	})
	emrCmd.AddCommand(emr_putManagedScalingPolicyCmd)
}
