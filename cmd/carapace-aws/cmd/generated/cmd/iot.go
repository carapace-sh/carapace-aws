package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotCmd = &cobra.Command{
	Use:   "iot",
	Short: "IoT\n\nIoT provides secure, bi-directional communication between Internet-connected devices (such as sensors, actuators, embedded devices, or smart appliances) and the Amazon Web Services cloud.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotCmd).Standalone()

	})
	rootCmd.AddCommand(iotCmd)
}
