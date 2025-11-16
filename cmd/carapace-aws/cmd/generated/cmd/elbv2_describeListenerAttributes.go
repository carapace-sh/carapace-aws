package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_describeListenerAttributesCmd = &cobra.Command{
	Use:   "describe-listener-attributes",
	Short: "Describes the attributes for the specified listener.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_describeListenerAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elbv2_describeListenerAttributesCmd).Standalone()

		elbv2_describeListenerAttributesCmd.Flags().String("listener-arn", "", "The Amazon Resource Name (ARN) of the listener.")
		elbv2_describeListenerAttributesCmd.MarkFlagRequired("listener-arn")
	})
	elbv2Cmd.AddCommand(elbv2_describeListenerAttributesCmd)
}
