package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updatePackageVersionCmd = &cobra.Command{
	Use:   "update-package-version",
	Short: "Updates the supported fields for a specific package version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updatePackageVersionCmd).Standalone()

	iot_updatePackageVersionCmd.Flags().String("action", "", "The status that the package version should be assigned.")
	iot_updatePackageVersionCmd.Flags().String("artifact", "", "The various components that make up a software package version.")
	iot_updatePackageVersionCmd.Flags().String("attributes", "", "Metadata that can be used to define a package version’s configuration.")
	iot_updatePackageVersionCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
	iot_updatePackageVersionCmd.Flags().String("description", "", "The package version description.")
	iot_updatePackageVersionCmd.Flags().String("package-name", "", "The name of the associated software package.")
	iot_updatePackageVersionCmd.Flags().String("recipe", "", "The inline job document associated with a software package version used for a quick job deployment.")
	iot_updatePackageVersionCmd.Flags().String("version-name", "", "The name of the target package version.")
	iot_updatePackageVersionCmd.MarkFlagRequired("package-name")
	iot_updatePackageVersionCmd.MarkFlagRequired("version-name")
	iotCmd.AddCommand(iot_updatePackageVersionCmd)
}
