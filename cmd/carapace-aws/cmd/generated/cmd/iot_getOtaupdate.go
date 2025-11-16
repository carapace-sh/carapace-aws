package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_getOtaupdateCmd = &cobra.Command{
	Use:   "get-otaupdate",
	Short: "Gets an OTA update.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_getOtaupdateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_getOtaupdateCmd).Standalone()

		iot_getOtaupdateCmd.Flags().String("ota-update-id", "", "The OTA update ID.")
		iot_getOtaupdateCmd.MarkFlagRequired("ota-update-id")
	})
	iotCmd.AddCommand(iot_getOtaupdateCmd)
}
