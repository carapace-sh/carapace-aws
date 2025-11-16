package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_getPackageVersionCmd = &cobra.Command{
	Use:   "get-package-version",
	Short: "Gets information about the specified package version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_getPackageVersionCmd).Standalone()

	iot_getPackageVersionCmd.Flags().String("package-name", "", "The name of the associated package.")
	iot_getPackageVersionCmd.Flags().String("version-name", "", "The name of the target package version.")
	iot_getPackageVersionCmd.MarkFlagRequired("package-name")
	iot_getPackageVersionCmd.MarkFlagRequired("version-name")
	iotCmd.AddCommand(iot_getPackageVersionCmd)
}
