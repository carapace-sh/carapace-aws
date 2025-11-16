package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updatePackageCmd = &cobra.Command{
	Use:   "update-package",
	Short: "Updates the supported fields for a specific software package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updatePackageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_updatePackageCmd).Standalone()

		iot_updatePackageCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
		iot_updatePackageCmd.Flags().String("default-version-name", "", "The name of the default package version.")
		iot_updatePackageCmd.Flags().String("description", "", "The package description.")
		iot_updatePackageCmd.Flags().String("package-name", "", "The name of the target software package.")
		iot_updatePackageCmd.Flags().String("unset-default-version", "", "Indicates whether you want to remove the named default package version from the software package.")
		iot_updatePackageCmd.MarkFlagRequired("package-name")
	})
	iotCmd.AddCommand(iot_updatePackageCmd)
}
