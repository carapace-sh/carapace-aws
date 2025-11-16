package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_modifyListenerAttributesCmd = &cobra.Command{
	Use:   "modify-listener-attributes",
	Short: "Modifies the specified attributes of the specified listener.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_modifyListenerAttributesCmd).Standalone()

	elbv2_modifyListenerAttributesCmd.Flags().String("attributes", "", "The listener attributes.")
	elbv2_modifyListenerAttributesCmd.Flags().String("listener-arn", "", "The Amazon Resource Name (ARN) of the listener.")
	elbv2_modifyListenerAttributesCmd.MarkFlagRequired("attributes")
	elbv2_modifyListenerAttributesCmd.MarkFlagRequired("listener-arn")
	elbv2Cmd.AddCommand(elbv2_modifyListenerAttributesCmd)
}
