package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotData_deleteConnectionCmd = &cobra.Command{
	Use:   "delete-connection",
	Short: "Disconnects a connected MQTT client from Amazon Web Services IoT Core.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotData_deleteConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotData_deleteConnectionCmd).Standalone()

		iotData_deleteConnectionCmd.Flags().String("clean-session", "", "Specifies whether to remove the client's session state when disconnecting.")
		iotData_deleteConnectionCmd.Flags().String("client-id", "", "The unique identifier of the MQTT client to disconnect.")
		iotData_deleteConnectionCmd.Flags().String("prevent-will-message", "", "Controls if Amazon Web Services IoT Core publishes the client's Last Will and Testament (LWT) message upon disconnection.")
		iotData_deleteConnectionCmd.MarkFlagRequired("client-id")
	})
	iotDataCmd.AddCommand(iotData_deleteConnectionCmd)
}
