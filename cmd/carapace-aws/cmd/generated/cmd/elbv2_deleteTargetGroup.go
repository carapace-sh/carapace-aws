package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_deleteTargetGroupCmd = &cobra.Command{
	Use:   "delete-target-group",
	Short: "Deletes the specified target group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_deleteTargetGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elbv2_deleteTargetGroupCmd).Standalone()

		elbv2_deleteTargetGroupCmd.Flags().String("target-group-arn", "", "The Amazon Resource Name (ARN) of the target group.")
		elbv2_deleteTargetGroupCmd.MarkFlagRequired("target-group-arn")
	})
	elbv2Cmd.AddCommand(elbv2_deleteTargetGroupCmd)
}
