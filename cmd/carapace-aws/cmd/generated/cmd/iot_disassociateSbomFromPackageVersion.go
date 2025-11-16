package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_disassociateSbomFromPackageVersionCmd = &cobra.Command{
	Use:   "disassociate-sbom-from-package-version",
	Short: "Disassociates the selected software bill of materials (SBOM) from a specific software package version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_disassociateSbomFromPackageVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_disassociateSbomFromPackageVersionCmd).Standalone()

		iot_disassociateSbomFromPackageVersionCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
		iot_disassociateSbomFromPackageVersionCmd.Flags().String("package-name", "", "The name of the new software package.")
		iot_disassociateSbomFromPackageVersionCmd.Flags().String("version-name", "", "The name of the new package version.")
		iot_disassociateSbomFromPackageVersionCmd.MarkFlagRequired("package-name")
		iot_disassociateSbomFromPackageVersionCmd.MarkFlagRequired("version-name")
	})
	iotCmd.AddCommand(iot_disassociateSbomFromPackageVersionCmd)
}
