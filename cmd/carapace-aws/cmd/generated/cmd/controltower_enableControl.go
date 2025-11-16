package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltower_enableControlCmd = &cobra.Command{
	Use:   "enable-control",
	Short: "This API call activates a control.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltower_enableControlCmd).Standalone()

	controltower_enableControlCmd.Flags().String("control-identifier", "", "The ARN of the control.")
	controltower_enableControlCmd.Flags().String("parameters", "", "A list of input parameter values, which are specified to configure the control when you enable it.")
	controltower_enableControlCmd.Flags().String("tags", "", "Tags to be applied to the `EnabledControl` resource.")
	controltower_enableControlCmd.Flags().String("target-identifier", "", "The ARN of the organizational unit.")
	controltower_enableControlCmd.MarkFlagRequired("control-identifier")
	controltower_enableControlCmd.MarkFlagRequired("target-identifier")
	controltowerCmd.AddCommand(controltower_enableControlCmd)
}
