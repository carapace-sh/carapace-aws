package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_registerTargetsCmd = &cobra.Command{
	Use:   "register-targets",
	Short: "Registers the specified targets with the specified target group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_registerTargetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elbv2_registerTargetsCmd).Standalone()

		elbv2_registerTargetsCmd.Flags().String("target-group-arn", "", "The Amazon Resource Name (ARN) of the target group.")
		elbv2_registerTargetsCmd.Flags().String("targets", "", "The targets.")
		elbv2_registerTargetsCmd.MarkFlagRequired("target-group-arn")
		elbv2_registerTargetsCmd.MarkFlagRequired("targets")
	})
	elbv2Cmd.AddCommand(elbv2_registerTargetsCmd)
}
