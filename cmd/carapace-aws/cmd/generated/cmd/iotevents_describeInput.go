package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotevents_describeInputCmd = &cobra.Command{
	Use:   "describe-input",
	Short: "Describes an input.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotevents_describeInputCmd).Standalone()

	iotevents_describeInputCmd.Flags().String("input-name", "", "The name of the input.")
	iotevents_describeInputCmd.MarkFlagRequired("input-name")
	ioteventsCmd.AddCommand(iotevents_describeInputCmd)
}
