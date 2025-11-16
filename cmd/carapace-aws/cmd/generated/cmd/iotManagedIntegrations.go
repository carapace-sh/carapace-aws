package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrationsCmd = &cobra.Command{
	Use:   "iot-managed-integrations",
	Short: "Managed integrations is a feature of AWS IoT Device Management that enables developers to quickly build innovative IoT solutions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrationsCmd).Standalone()

	rootCmd.AddCommand(iotManagedIntegrationsCmd)
}
