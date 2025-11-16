package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listOtaupdatesCmd = &cobra.Command{
	Use:   "list-otaupdates",
	Short: "Lists OTA updates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listOtaupdatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listOtaupdatesCmd).Standalone()

		iot_listOtaupdatesCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
		iot_listOtaupdatesCmd.Flags().String("next-token", "", "A token used to retrieve the next set of results.")
		iot_listOtaupdatesCmd.Flags().String("ota-update-status", "", "The OTA update job status.")
	})
	iotCmd.AddCommand(iot_listOtaupdatesCmd)
}
