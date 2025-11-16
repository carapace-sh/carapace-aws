package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ioteventsData_batchDeleteDetectorCmd = &cobra.Command{
	Use:   "batch-delete-detector",
	Short: "Deletes one or more detectors that were created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ioteventsData_batchDeleteDetectorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ioteventsData_batchDeleteDetectorCmd).Standalone()

		ioteventsData_batchDeleteDetectorCmd.Flags().String("detectors", "", "The list of one or more detectors to be deleted.")
		ioteventsData_batchDeleteDetectorCmd.MarkFlagRequired("detectors")
	})
	ioteventsDataCmd.AddCommand(ioteventsData_batchDeleteDetectorCmd)
}
