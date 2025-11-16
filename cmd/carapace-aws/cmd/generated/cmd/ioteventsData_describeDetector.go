package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ioteventsData_describeDetectorCmd = &cobra.Command{
	Use:   "describe-detector",
	Short: "Returns information about the specified detector (instance).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ioteventsData_describeDetectorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ioteventsData_describeDetectorCmd).Standalone()

		ioteventsData_describeDetectorCmd.Flags().String("detector-model-name", "", "The name of the detector model whose detectors (instances) you want information about.")
		ioteventsData_describeDetectorCmd.Flags().String("key-value", "", "A filter used to limit results to detectors (instances) created because of the given key ID.")
		ioteventsData_describeDetectorCmd.MarkFlagRequired("detector-model-name")
	})
	ioteventsDataCmd.AddCommand(ioteventsData_describeDetectorCmd)
}
