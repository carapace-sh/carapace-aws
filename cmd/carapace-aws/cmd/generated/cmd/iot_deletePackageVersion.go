package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deletePackageVersionCmd = &cobra.Command{
	Use:   "delete-package-version",
	Short: "Deletes a specific version from a software package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deletePackageVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_deletePackageVersionCmd).Standalone()

		iot_deletePackageVersionCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
		iot_deletePackageVersionCmd.Flags().String("package-name", "", "The name of the associated software package.")
		iot_deletePackageVersionCmd.Flags().String("version-name", "", "The name of the target package version.")
		iot_deletePackageVersionCmd.MarkFlagRequired("package-name")
		iot_deletePackageVersionCmd.MarkFlagRequired("version-name")
	})
	iotCmd.AddCommand(iot_deletePackageVersionCmd)
}
