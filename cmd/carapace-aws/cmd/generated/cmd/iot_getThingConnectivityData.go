package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_getThingConnectivityDataCmd = &cobra.Command{
	Use:   "get-thing-connectivity-data",
	Short: "Retrieves the live connectivity status per device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_getThingConnectivityDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_getThingConnectivityDataCmd).Standalone()

		iot_getThingConnectivityDataCmd.Flags().String("thing-name", "", "The name of your IoT thing.")
		iot_getThingConnectivityDataCmd.MarkFlagRequired("thing-name")
	})
	iotCmd.AddCommand(iot_getThingConnectivityDataCmd)
}
