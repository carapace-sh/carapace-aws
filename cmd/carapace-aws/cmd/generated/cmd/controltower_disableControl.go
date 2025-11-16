package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltower_disableControlCmd = &cobra.Command{
	Use:   "disable-control",
	Short: "This API call turns off a control.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltower_disableControlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(controltower_disableControlCmd).Standalone()

		controltower_disableControlCmd.Flags().String("control-identifier", "", "The ARN of the control.")
		controltower_disableControlCmd.Flags().String("enabled-control-identifier", "", "The ARN of the enabled control to be disabled, which uniquely identifies the control instance on the target organizational unit.")
		controltower_disableControlCmd.Flags().String("target-identifier", "", "The ARN of the organizational unit.")
	})
	controltowerCmd.AddCommand(controltower_disableControlCmd)
}
