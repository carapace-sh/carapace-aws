package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_getPackageVersionAssetCmd = &cobra.Command{
	Use:   "get-package-version-asset",
	Short: "Returns an asset (or file) that is in a package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_getPackageVersionAssetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeartifact_getPackageVersionAssetCmd).Standalone()

		codeartifact_getPackageVersionAssetCmd.Flags().String("asset", "", "The name of the requested asset.")
		codeartifact_getPackageVersionAssetCmd.Flags().String("domain", "", "The name of the domain that contains the repository that contains the package version with the requested asset.")
		codeartifact_getPackageVersionAssetCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
		codeartifact_getPackageVersionAssetCmd.Flags().String("format", "", "A format that specifies the type of the package version with the requested asset file.")
		codeartifact_getPackageVersionAssetCmd.Flags().String("namespace", "", "The namespace of the package version with the requested asset file.")
		codeartifact_getPackageVersionAssetCmd.Flags().String("package", "", "The name of the package that contains the requested asset.")
		codeartifact_getPackageVersionAssetCmd.Flags().String("package-version", "", "A string that contains the package version (for example, `3.5.2`).")
		codeartifact_getPackageVersionAssetCmd.Flags().String("package-version-revision", "", "The name of the package version revision that contains the requested asset.")
		codeartifact_getPackageVersionAssetCmd.Flags().String("repository", "", "The repository that contains the package version with the requested asset.")
		codeartifact_getPackageVersionAssetCmd.MarkFlagRequired("asset")
		codeartifact_getPackageVersionAssetCmd.MarkFlagRequired("domain")
		codeartifact_getPackageVersionAssetCmd.MarkFlagRequired("format")
		codeartifact_getPackageVersionAssetCmd.MarkFlagRequired("package")
		codeartifact_getPackageVersionAssetCmd.MarkFlagRequired("package-version")
		codeartifact_getPackageVersionAssetCmd.MarkFlagRequired("repository")
	})
	codeartifactCmd.AddCommand(codeartifact_getPackageVersionAssetCmd)
}
