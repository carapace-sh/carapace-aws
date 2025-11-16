package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_deregisterTargetsCmd = &cobra.Command{
	Use:   "deregister-targets",
	Short: "Deregisters the specified targets from the specified target group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_deregisterTargetsCmd).Standalone()

	elbv2_deregisterTargetsCmd.Flags().String("target-group-arn", "", "The Amazon Resource Name (ARN) of the target group.")
	elbv2_deregisterTargetsCmd.Flags().String("targets", "", "The targets.")
	elbv2_deregisterTargetsCmd.MarkFlagRequired("target-group-arn")
	elbv2_deregisterTargetsCmd.MarkFlagRequired("targets")
	elbv2Cmd.AddCommand(elbv2_deregisterTargetsCmd)
}
