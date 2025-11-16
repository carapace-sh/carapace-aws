package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotevents_updateInputCmd = &cobra.Command{
	Use:   "update-input",
	Short: "Updates an input.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotevents_updateInputCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotevents_updateInputCmd).Standalone()

		iotevents_updateInputCmd.Flags().String("input-definition", "", "The definition of the input.")
		iotevents_updateInputCmd.Flags().String("input-description", "", "A brief description of the input.")
		iotevents_updateInputCmd.Flags().String("input-name", "", "The name of the input you want to update.")
		iotevents_updateInputCmd.MarkFlagRequired("input-definition")
		iotevents_updateInputCmd.MarkFlagRequired("input-name")
	})
	ioteventsCmd.AddCommand(iotevents_updateInputCmd)
}
