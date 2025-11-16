package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_updatePackageVersionsStatusCmd = &cobra.Command{
	Use:   "update-package-versions-status",
	Short: "Updates the status of one or more versions of a package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_updatePackageVersionsStatusCmd).Standalone()

	codeartifact_updatePackageVersionsStatusCmd.Flags().String("domain", "", "The name of the domain that contains the repository that contains the package versions with a status to be updated.")
	codeartifact_updatePackageVersionsStatusCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
	codeartifact_updatePackageVersionsStatusCmd.Flags().String("expected-status", "", "The package version’s expected status before it is updated.")
	codeartifact_updatePackageVersionsStatusCmd.Flags().String("format", "", "A format that specifies the type of the package with the statuses to update.")
	codeartifact_updatePackageVersionsStatusCmd.Flags().String("namespace", "", "The namespace of the package version to be updated.")
	codeartifact_updatePackageVersionsStatusCmd.Flags().String("package", "", "The name of the package with the version statuses to update.")
	codeartifact_updatePackageVersionsStatusCmd.Flags().String("repository", "", "The repository that contains the package versions with the status you want to update.")
	codeartifact_updatePackageVersionsStatusCmd.Flags().String("target-status", "", "The status you want to change the package version status to.")
	codeartifact_updatePackageVersionsStatusCmd.Flags().String("version-revisions", "", "A map of package versions and package version revisions.")
	codeartifact_updatePackageVersionsStatusCmd.Flags().String("versions", "", "An array of strings that specify the versions of the package with the statuses to update.")
	codeartifact_updatePackageVersionsStatusCmd.MarkFlagRequired("domain")
	codeartifact_updatePackageVersionsStatusCmd.MarkFlagRequired("format")
	codeartifact_updatePackageVersionsStatusCmd.MarkFlagRequired("package")
	codeartifact_updatePackageVersionsStatusCmd.MarkFlagRequired("repository")
	codeartifact_updatePackageVersionsStatusCmd.MarkFlagRequired("target-status")
	codeartifact_updatePackageVersionsStatusCmd.MarkFlagRequired("versions")
	codeartifactCmd.AddCommand(codeartifact_updatePackageVersionsStatusCmd)
}
