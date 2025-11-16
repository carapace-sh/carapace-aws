package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ioteventsData_batchUpdateDetectorCmd = &cobra.Command{
	Use:   "batch-update-detector",
	Short: "Updates the state, variable values, and timer settings of one or more detectors (instances) of a specified detector model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ioteventsData_batchUpdateDetectorCmd).Standalone()

	ioteventsData_batchUpdateDetectorCmd.Flags().String("detectors", "", "The list of detectors (instances) to update, along with the values to update.")
	ioteventsData_batchUpdateDetectorCmd.MarkFlagRequired("detectors")
	ioteventsDataCmd.AddCommand(ioteventsData_batchUpdateDetectorCmd)
}
