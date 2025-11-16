package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_modifyTargetGroupAttributesCmd = &cobra.Command{
	Use:   "modify-target-group-attributes",
	Short: "Modifies the specified attributes of the specified target group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_modifyTargetGroupAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elbv2_modifyTargetGroupAttributesCmd).Standalone()

		elbv2_modifyTargetGroupAttributesCmd.Flags().String("attributes", "", "The target group attributes.")
		elbv2_modifyTargetGroupAttributesCmd.Flags().String("target-group-arn", "", "The Amazon Resource Name (ARN) of the target group.")
		elbv2_modifyTargetGroupAttributesCmd.MarkFlagRequired("attributes")
		elbv2_modifyTargetGroupAttributesCmd.MarkFlagRequired("target-group-arn")
	})
	elbv2Cmd.AddCommand(elbv2_modifyTargetGroupAttributesCmd)
}
