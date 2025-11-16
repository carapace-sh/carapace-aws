package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_getAutoTerminationPolicyCmd = &cobra.Command{
	Use:   "get-auto-termination-policy",
	Short: "Returns the auto-termination policy for an Amazon EMR cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_getAutoTerminationPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_getAutoTerminationPolicyCmd).Standalone()

		emr_getAutoTerminationPolicyCmd.Flags().String("cluster-id", "", "Specifies the ID of the Amazon EMR cluster for which the auto-termination policy will be fetched.")
		emr_getAutoTerminationPolicyCmd.MarkFlagRequired("cluster-id")
	})
	emrCmd.AddCommand(emr_getAutoTerminationPolicyCmd)
}
