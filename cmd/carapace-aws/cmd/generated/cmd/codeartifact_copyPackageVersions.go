package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_copyPackageVersionsCmd = &cobra.Command{
	Use:   "copy-package-versions",
	Short: "Copies package versions from one repository to another repository in the same domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_copyPackageVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeartifact_copyPackageVersionsCmd).Standalone()

		codeartifact_copyPackageVersionsCmd.Flags().String("allow-overwrite", "", "Set to true to overwrite a package version that already exists in the destination repository.")
		codeartifact_copyPackageVersionsCmd.Flags().String("destination-repository", "", "The name of the repository into which package versions are copied.")
		codeartifact_copyPackageVersionsCmd.Flags().String("domain", "", "The name of the domain that contains the source and destination repositories.")
		codeartifact_copyPackageVersionsCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
		codeartifact_copyPackageVersionsCmd.Flags().String("format", "", "The format of the package versions to be copied.")
		codeartifact_copyPackageVersionsCmd.Flags().String("include-from-upstream", "", "Set to true to copy packages from repositories that are upstream from the source repository to the destination repository.")
		codeartifact_copyPackageVersionsCmd.Flags().String("namespace", "", "The namespace of the package versions to be copied.")
		codeartifact_copyPackageVersionsCmd.Flags().String("package", "", "The name of the package that contains the versions to be copied.")
		codeartifact_copyPackageVersionsCmd.Flags().String("source-repository", "", "The name of the repository that contains the package versions to be copied.")
		codeartifact_copyPackageVersionsCmd.Flags().String("version-revisions", "", "A list of key-value pairs.")
		codeartifact_copyPackageVersionsCmd.Flags().String("versions", "", "The versions of the package to be copied.")
		codeartifact_copyPackageVersionsCmd.MarkFlagRequired("destination-repository")
		codeartifact_copyPackageVersionsCmd.MarkFlagRequired("domain")
		codeartifact_copyPackageVersionsCmd.MarkFlagRequired("format")
		codeartifact_copyPackageVersionsCmd.MarkFlagRequired("package")
		codeartifact_copyPackageVersionsCmd.MarkFlagRequired("source-repository")
	})
	codeartifactCmd.AddCommand(codeartifact_copyPackageVersionsCmd)
}
