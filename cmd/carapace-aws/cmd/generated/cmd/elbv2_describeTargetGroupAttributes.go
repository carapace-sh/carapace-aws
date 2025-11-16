package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_describeTargetGroupAttributesCmd = &cobra.Command{
	Use:   "describe-target-group-attributes",
	Short: "Describes the attributes for the specified target group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_describeTargetGroupAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elbv2_describeTargetGroupAttributesCmd).Standalone()

		elbv2_describeTargetGroupAttributesCmd.Flags().String("target-group-arn", "", "The Amazon Resource Name (ARN) of the target group.")
		elbv2_describeTargetGroupAttributesCmd.MarkFlagRequired("target-group-arn")
	})
	elbv2Cmd.AddCommand(elbv2_describeTargetGroupAttributesCmd)
}
