package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_listInputDevicesCmd = &cobra.Command{
	Use:   "list-input-devices",
	Short: "List input devices",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_listInputDevicesCmd).Standalone()

	medialive_listInputDevicesCmd.Flags().String("max-results", "", "")
	medialive_listInputDevicesCmd.Flags().String("next-token", "", "")
	medialiveCmd.AddCommand(medialive_listInputDevicesCmd)
}
