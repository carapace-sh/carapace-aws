package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteOtaupdateCmd = &cobra.Command{
	Use:   "delete-otaupdate",
	Short: "Delete an OTA update.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteOtaupdateCmd).Standalone()

	iot_deleteOtaupdateCmd.Flags().String("delete-stream", "", "When true, the stream created by the OTAUpdate process is deleted when the OTA update is deleted.")
	iot_deleteOtaupdateCmd.Flags().String("force-delete-awsjob", "", "When true, deletes the IoT job created by the OTAUpdate process even if it is \"IN\\_PROGRESS\".")
	iot_deleteOtaupdateCmd.Flags().String("ota-update-id", "", "The ID of the OTA update to delete.")
	iot_deleteOtaupdateCmd.MarkFlagRequired("ota-update-id")
	iotCmd.AddCommand(iot_deleteOtaupdateCmd)
}
