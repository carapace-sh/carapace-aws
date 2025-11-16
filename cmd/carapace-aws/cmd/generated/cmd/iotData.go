package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotDataCmd = &cobra.Command{
	Use:   "iot-data",
	Short: "IoT data\n\nIoT data enables secure, bi-directional communication between Internet-connected things (such as sensors, actuators, embedded devices, or smart appliances) and the Amazon Web Services cloud.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotDataCmd).Standalone()

	rootCmd.AddCommand(iotDataCmd)
}
