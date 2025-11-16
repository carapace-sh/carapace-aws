package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_describeTargetHealthCmd = &cobra.Command{
	Use:   "describe-target-health",
	Short: "Describes the health of the specified targets or all of your targets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_describeTargetHealthCmd).Standalone()

	elbv2_describeTargetHealthCmd.Flags().String("include", "", "Used to include anomaly detection information.")
	elbv2_describeTargetHealthCmd.Flags().String("target-group-arn", "", "The Amazon Resource Name (ARN) of the target group.")
	elbv2_describeTargetHealthCmd.Flags().String("targets", "", "The targets.")
	elbv2_describeTargetHealthCmd.MarkFlagRequired("target-group-arn")
	elbv2Cmd.AddCommand(elbv2_describeTargetHealthCmd)
}
