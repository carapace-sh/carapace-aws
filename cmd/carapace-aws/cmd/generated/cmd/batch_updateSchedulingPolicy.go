package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_updateSchedulingPolicyCmd = &cobra.Command{
	Use:   "update-scheduling-policy",
	Short: "Updates a scheduling policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_updateSchedulingPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(batch_updateSchedulingPolicyCmd).Standalone()

		batch_updateSchedulingPolicyCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the scheduling policy to update.")
		batch_updateSchedulingPolicyCmd.Flags().String("fairshare-policy", "", "The fair-share policy scheduling details.")
		batch_updateSchedulingPolicyCmd.MarkFlagRequired("arn")
	})
	batchCmd.AddCommand(batch_updateSchedulingPolicyCmd)
}
