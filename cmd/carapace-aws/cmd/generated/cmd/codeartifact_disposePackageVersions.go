package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_disposePackageVersionsCmd = &cobra.Command{
	Use:   "dispose-package-versions",
	Short: "Deletes the assets in package versions and sets the package versions' status to `Disposed`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_disposePackageVersionsCmd).Standalone()

	codeartifact_disposePackageVersionsCmd.Flags().String("domain", "", "The name of the domain that contains the repository you want to dispose.")
	codeartifact_disposePackageVersionsCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
	codeartifact_disposePackageVersionsCmd.Flags().String("expected-status", "", "The expected status of the package version to dispose.")
	codeartifact_disposePackageVersionsCmd.Flags().String("format", "", "A format that specifies the type of package versions you want to dispose.")
	codeartifact_disposePackageVersionsCmd.Flags().String("namespace", "", "The namespace of the package versions to be disposed.")
	codeartifact_disposePackageVersionsCmd.Flags().String("package", "", "The name of the package with the versions you want to dispose.")
	codeartifact_disposePackageVersionsCmd.Flags().String("repository", "", "The name of the repository that contains the package versions you want to dispose.")
	codeartifact_disposePackageVersionsCmd.Flags().String("version-revisions", "", "The revisions of the package versions you want to dispose.")
	codeartifact_disposePackageVersionsCmd.Flags().String("versions", "", "The versions of the package you want to dispose.")
	codeartifact_disposePackageVersionsCmd.MarkFlagRequired("domain")
	codeartifact_disposePackageVersionsCmd.MarkFlagRequired("format")
	codeartifact_disposePackageVersionsCmd.MarkFlagRequired("package")
	codeartifact_disposePackageVersionsCmd.MarkFlagRequired("repository")
	codeartifact_disposePackageVersionsCmd.MarkFlagRequired("versions")
	codeartifactCmd.AddCommand(codeartifact_disposePackageVersionsCmd)
}
