package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotevents_createInputCmd = &cobra.Command{
	Use:   "create-input",
	Short: "Creates an input.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotevents_createInputCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotevents_createInputCmd).Standalone()

		iotevents_createInputCmd.Flags().String("input-definition", "", "The definition of the input.")
		iotevents_createInputCmd.Flags().String("input-description", "", "A brief description of the input.")
		iotevents_createInputCmd.Flags().String("input-name", "", "The name you want to give to the input.")
		iotevents_createInputCmd.Flags().String("tags", "", "Metadata that can be used to manage the input.")
		iotevents_createInputCmd.MarkFlagRequired("input-definition")
		iotevents_createInputCmd.MarkFlagRequired("input-name")
	})
	ioteventsCmd.AddCommand(iotevents_createInputCmd)
}
