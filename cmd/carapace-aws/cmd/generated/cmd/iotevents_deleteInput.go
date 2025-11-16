package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotevents_deleteInputCmd = &cobra.Command{
	Use:   "delete-input",
	Short: "Deletes an input.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotevents_deleteInputCmd).Standalone()

	iotevents_deleteInputCmd.Flags().String("input-name", "", "The name of the input to delete.")
	iotevents_deleteInputCmd.MarkFlagRequired("input-name")
	ioteventsCmd.AddCommand(iotevents_deleteInputCmd)
}
