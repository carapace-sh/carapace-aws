package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_deleteSchedulingPolicyCmd = &cobra.Command{
	Use:   "delete-scheduling-policy",
	Short: "Deletes the specified scheduling policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_deleteSchedulingPolicyCmd).Standalone()

	batch_deleteSchedulingPolicyCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the scheduling policy to delete.")
	batch_deleteSchedulingPolicyCmd.MarkFlagRequired("arn")
	batchCmd.AddCommand(batch_deleteSchedulingPolicyCmd)
}
