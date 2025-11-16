package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotdeviceadvisorCmd = &cobra.Command{
	Use:   "iotdeviceadvisor",
	Short: "Amazon Web Services IoT Core Device Advisor is a cloud-based, fully managed test capability for validating IoT devices during device software development.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotdeviceadvisorCmd).Standalone()

	rootCmd.AddCommand(iotdeviceadvisorCmd)
}
