package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_listPackageVersionDependenciesCmd = &cobra.Command{
	Use:   "list-package-version-dependencies",
	Short: "Returns the direct dependencies for a package version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_listPackageVersionDependenciesCmd).Standalone()

	codeartifact_listPackageVersionDependenciesCmd.Flags().String("domain", "", "The name of the domain that contains the repository that contains the requested package version dependencies.")
	codeartifact_listPackageVersionDependenciesCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
	codeartifact_listPackageVersionDependenciesCmd.Flags().String("format", "", "The format of the package with the requested dependencies.")
	codeartifact_listPackageVersionDependenciesCmd.Flags().String("namespace", "", "The namespace of the package version with the requested dependencies.")
	codeartifact_listPackageVersionDependenciesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	codeartifact_listPackageVersionDependenciesCmd.Flags().String("package", "", "The name of the package versions' package.")
	codeartifact_listPackageVersionDependenciesCmd.Flags().String("package-version", "", "A string that contains the package version (for example, `3.5.2`).")
	codeartifact_listPackageVersionDependenciesCmd.Flags().String("repository", "", "The name of the repository that contains the requested package version.")
	codeartifact_listPackageVersionDependenciesCmd.MarkFlagRequired("domain")
	codeartifact_listPackageVersionDependenciesCmd.MarkFlagRequired("format")
	codeartifact_listPackageVersionDependenciesCmd.MarkFlagRequired("package")
	codeartifact_listPackageVersionDependenciesCmd.MarkFlagRequired("package-version")
	codeartifact_listPackageVersionDependenciesCmd.MarkFlagRequired("repository")
	codeartifactCmd.AddCommand(codeartifact_listPackageVersionDependenciesCmd)
}
