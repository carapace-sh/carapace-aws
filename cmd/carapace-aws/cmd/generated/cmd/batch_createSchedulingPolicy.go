package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_createSchedulingPolicyCmd = &cobra.Command{
	Use:   "create-scheduling-policy",
	Short: "Creates an Batch scheduling policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_createSchedulingPolicyCmd).Standalone()

	batch_createSchedulingPolicyCmd.Flags().String("fairshare-policy", "", "The fair-share scheduling policy details.")
	batch_createSchedulingPolicyCmd.Flags().String("name", "", "The name of the fair-share scheduling policy.")
	batch_createSchedulingPolicyCmd.Flags().String("tags", "", "The tags that you apply to the scheduling policy to help you categorize and organize your resources.")
	batch_createSchedulingPolicyCmd.MarkFlagRequired("name")
	batchCmd.AddCommand(batch_createSchedulingPolicyCmd)
}
