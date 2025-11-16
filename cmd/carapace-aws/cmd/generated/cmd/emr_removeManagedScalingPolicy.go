package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_removeManagedScalingPolicyCmd = &cobra.Command{
	Use:   "remove-managed-scaling-policy",
	Short: "Removes a managed scaling policy from a specified Amazon EMR cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_removeManagedScalingPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_removeManagedScalingPolicyCmd).Standalone()

		emr_removeManagedScalingPolicyCmd.Flags().String("cluster-id", "", "Specifies the ID of the cluster from which the managed scaling policy will be removed.")
		emr_removeManagedScalingPolicyCmd.MarkFlagRequired("cluster-id")
	})
	emrCmd.AddCommand(emr_removeManagedScalingPolicyCmd)
}
