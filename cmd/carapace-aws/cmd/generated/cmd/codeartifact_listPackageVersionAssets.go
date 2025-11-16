package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_listPackageVersionAssetsCmd = &cobra.Command{
	Use:   "list-package-version-assets",
	Short: "Returns a list of [AssetSummary](https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_AssetSummary.html) objects for assets in a package version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_listPackageVersionAssetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeartifact_listPackageVersionAssetsCmd).Standalone()

		codeartifact_listPackageVersionAssetsCmd.Flags().String("domain", "", "The name of the domain that contains the repository associated with the package version assets.")
		codeartifact_listPackageVersionAssetsCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
		codeartifact_listPackageVersionAssetsCmd.Flags().String("format", "", "The format of the package that contains the requested package version assets.")
		codeartifact_listPackageVersionAssetsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		codeartifact_listPackageVersionAssetsCmd.Flags().String("namespace", "", "The namespace of the package version that contains the requested package version assets.")
		codeartifact_listPackageVersionAssetsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		codeartifact_listPackageVersionAssetsCmd.Flags().String("package", "", "The name of the package that contains the requested package version assets.")
		codeartifact_listPackageVersionAssetsCmd.Flags().String("package-version", "", "A string that contains the package version (for example, `3.5.2`).")
		codeartifact_listPackageVersionAssetsCmd.Flags().String("repository", "", "The name of the repository that contains the package that contains the requested package version assets.")
		codeartifact_listPackageVersionAssetsCmd.MarkFlagRequired("domain")
		codeartifact_listPackageVersionAssetsCmd.MarkFlagRequired("format")
		codeartifact_listPackageVersionAssetsCmd.MarkFlagRequired("package")
		codeartifact_listPackageVersionAssetsCmd.MarkFlagRequired("package-version")
		codeartifact_listPackageVersionAssetsCmd.MarkFlagRequired("repository")
	})
	codeartifactCmd.AddCommand(codeartifact_listPackageVersionAssetsCmd)
}
