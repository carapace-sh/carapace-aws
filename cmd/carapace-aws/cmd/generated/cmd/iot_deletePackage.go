package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deletePackageCmd = &cobra.Command{
	Use:   "delete-package",
	Short: "Deletes a specific version from a software package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deletePackageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_deletePackageCmd).Standalone()

		iot_deletePackageCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
		iot_deletePackageCmd.Flags().String("package-name", "", "The name of the target software package.")
		iot_deletePackageCmd.MarkFlagRequired("package-name")
	})
	iotCmd.AddCommand(iot_deletePackageCmd)
}
