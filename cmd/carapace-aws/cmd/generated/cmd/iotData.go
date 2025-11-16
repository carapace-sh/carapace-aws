package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotDataCmd = &cobra.Command{
	Use:   "iot-data",
	Short: "IoT data",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotDataCmd).Standalone()

	})
	rootCmd.AddCommand(iotDataCmd)
}
