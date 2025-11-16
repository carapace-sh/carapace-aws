package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_associateSbomWithPackageVersionCmd = &cobra.Command{
	Use:   "associate-sbom-with-package-version",
	Short: "Associates the selected software bill of materials (SBOM) with a specific software package version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_associateSbomWithPackageVersionCmd).Standalone()

	iot_associateSbomWithPackageVersionCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
	iot_associateSbomWithPackageVersionCmd.Flags().String("package-name", "", "The name of the new software package.")
	iot_associateSbomWithPackageVersionCmd.Flags().String("sbom", "", "")
	iot_associateSbomWithPackageVersionCmd.Flags().String("version-name", "", "The name of the new package version.")
	iot_associateSbomWithPackageVersionCmd.MarkFlagRequired("package-name")
	iot_associateSbomWithPackageVersionCmd.MarkFlagRequired("sbom")
	iot_associateSbomWithPackageVersionCmd.MarkFlagRequired("version-name")
	iotCmd.AddCommand(iot_associateSbomWithPackageVersionCmd)
}
