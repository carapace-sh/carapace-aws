package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_getManagedScalingPolicyCmd = &cobra.Command{
	Use:   "get-managed-scaling-policy",
	Short: "Fetches the attached managed scaling policy for an Amazon EMR cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_getManagedScalingPolicyCmd).Standalone()

	emr_getManagedScalingPolicyCmd.Flags().String("cluster-id", "", "Specifies the ID of the cluster for which the managed scaling policy will be fetched.")
	emr_getManagedScalingPolicyCmd.MarkFlagRequired("cluster-id")
	emrCmd.AddCommand(emr_getManagedScalingPolicyCmd)
}
