package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_deletePackageVersionsCmd = &cobra.Command{
	Use:   "delete-package-versions",
	Short: "Deletes one or more versions of a package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_deletePackageVersionsCmd).Standalone()

	codeartifact_deletePackageVersionsCmd.Flags().String("domain", "", "The name of the domain that contains the package to delete.")
	codeartifact_deletePackageVersionsCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
	codeartifact_deletePackageVersionsCmd.Flags().String("expected-status", "", "The expected status of the package version to delete.")
	codeartifact_deletePackageVersionsCmd.Flags().String("format", "", "The format of the package versions to delete.")
	codeartifact_deletePackageVersionsCmd.Flags().String("namespace", "", "The namespace of the package versions to be deleted.")
	codeartifact_deletePackageVersionsCmd.Flags().String("package", "", "The name of the package with the versions to delete.")
	codeartifact_deletePackageVersionsCmd.Flags().String("repository", "", "The name of the repository that contains the package versions to delete.")
	codeartifact_deletePackageVersionsCmd.Flags().String("versions", "", "An array of strings that specify the versions of the package to delete.")
	codeartifact_deletePackageVersionsCmd.MarkFlagRequired("domain")
	codeartifact_deletePackageVersionsCmd.MarkFlagRequired("format")
	codeartifact_deletePackageVersionsCmd.MarkFlagRequired("package")
	codeartifact_deletePackageVersionsCmd.MarkFlagRequired("repository")
	codeartifact_deletePackageVersionsCmd.MarkFlagRequired("versions")
	codeartifactCmd.AddCommand(codeartifact_deletePackageVersionsCmd)
}
