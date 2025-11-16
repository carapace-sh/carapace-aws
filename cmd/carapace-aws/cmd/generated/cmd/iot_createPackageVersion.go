package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createPackageVersionCmd = &cobra.Command{
	Use:   "create-package-version",
	Short: "Creates a new version for an existing IoT software package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createPackageVersionCmd).Standalone()

	iot_createPackageVersionCmd.Flags().String("artifact", "", "The various build components created during the build process such as libraries and configuration files that make up a software package version.")
	iot_createPackageVersionCmd.Flags().String("attributes", "", "Metadata that can be used to define a package version’s configuration.")
	iot_createPackageVersionCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
	iot_createPackageVersionCmd.Flags().String("description", "", "A summary of the package version being created.")
	iot_createPackageVersionCmd.Flags().String("package-name", "", "The name of the associated software package.")
	iot_createPackageVersionCmd.Flags().String("recipe", "", "The inline job document associated with a software package version used for a quick job deployment.")
	iot_createPackageVersionCmd.Flags().String("tags", "", "Metadata that can be used to manage the package version.")
	iot_createPackageVersionCmd.Flags().String("version-name", "", "The name of the new package version.")
	iot_createPackageVersionCmd.MarkFlagRequired("package-name")
	iot_createPackageVersionCmd.MarkFlagRequired("version-name")
	iotCmd.AddCommand(iot_createPackageVersionCmd)
}
