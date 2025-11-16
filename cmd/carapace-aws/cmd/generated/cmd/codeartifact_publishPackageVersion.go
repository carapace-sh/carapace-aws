package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_publishPackageVersionCmd = &cobra.Command{
	Use:   "publish-package-version",
	Short: "Creates a new package version containing one or more assets (or files).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_publishPackageVersionCmd).Standalone()

	codeartifact_publishPackageVersionCmd.Flags().String("asset-content", "", "The content of the asset to publish.")
	codeartifact_publishPackageVersionCmd.Flags().String("asset-name", "", "The name of the asset to publish.")
	codeartifact_publishPackageVersionCmd.Flags().String("asset-sha256", "", "The SHA256 hash of the `assetContent` to publish.")
	codeartifact_publishPackageVersionCmd.Flags().String("domain", "", "The name of the domain that contains the repository that contains the package version to publish.")
	codeartifact_publishPackageVersionCmd.Flags().String("domain-owner", "", "The 12-digit account number of the AWS account that owns the domain.")
	codeartifact_publishPackageVersionCmd.Flags().String("format", "", "A format that specifies the type of the package version with the requested asset file.")
	codeartifact_publishPackageVersionCmd.Flags().String("namespace", "", "The namespace of the package version to publish.")
	codeartifact_publishPackageVersionCmd.Flags().String("package", "", "The name of the package version to publish.")
	codeartifact_publishPackageVersionCmd.Flags().String("package-version", "", "The package version to publish (for example, `3.5.2`).")
	codeartifact_publishPackageVersionCmd.Flags().String("repository", "", "The name of the repository that the package version will be published to.")
	codeartifact_publishPackageVersionCmd.Flags().String("unfinished", "", "Specifies whether the package version should remain in the `unfinished` state.")
	codeartifact_publishPackageVersionCmd.MarkFlagRequired("asset-content")
	codeartifact_publishPackageVersionCmd.MarkFlagRequired("asset-name")
	codeartifact_publishPackageVersionCmd.MarkFlagRequired("asset-sha256")
	codeartifact_publishPackageVersionCmd.MarkFlagRequired("domain")
	codeartifact_publishPackageVersionCmd.MarkFlagRequired("format")
	codeartifact_publishPackageVersionCmd.MarkFlagRequired("package")
	codeartifact_publishPackageVersionCmd.MarkFlagRequired("package-version")
	codeartifact_publishPackageVersionCmd.MarkFlagRequired("repository")
	codeartifactCmd.AddCommand(codeartifact_publishPackageVersionCmd)
}
