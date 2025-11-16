package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_removeAutoTerminationPolicyCmd = &cobra.Command{
	Use:   "remove-auto-termination-policy",
	Short: "Removes an auto-termination policy from an Amazon EMR cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_removeAutoTerminationPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_removeAutoTerminationPolicyCmd).Standalone()

		emr_removeAutoTerminationPolicyCmd.Flags().String("cluster-id", "", "Specifies the ID of the Amazon EMR cluster from which the auto-termination policy will be removed.")
		emr_removeAutoTerminationPolicyCmd.MarkFlagRequired("cluster-id")
	})
	emrCmd.AddCommand(emr_removeAutoTerminationPolicyCmd)
}
