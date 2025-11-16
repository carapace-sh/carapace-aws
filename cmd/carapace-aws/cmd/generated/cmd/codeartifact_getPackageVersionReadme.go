package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_getPackageVersionReadmeCmd = &cobra.Command{
	Use:   "get-package-version-readme",
	Short: "Gets the readme file or descriptive text for a package version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_getPackageVersionReadmeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeartifact_getPackageVersionReadmeCmd).Standalone()

		codeartifact_getPackageVersionReadmeCmd.Flags().String("domain", "", "The name of the domain that contains the repository that contains the package version with the requested readme file.")
		codeartifact_getPackageVersionReadmeCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
		codeartifact_getPackageVersionReadmeCmd.Flags().String("format", "", "A format that specifies the type of the package version with the requested readme file.")
		codeartifact_getPackageVersionReadmeCmd.Flags().String("namespace", "", "The namespace of the package version with the requested readme file.")
		codeartifact_getPackageVersionReadmeCmd.Flags().String("package", "", "The name of the package version that contains the requested readme file.")
		codeartifact_getPackageVersionReadmeCmd.Flags().String("package-version", "", "A string that contains the package version (for example, `3.5.2`).")
		codeartifact_getPackageVersionReadmeCmd.Flags().String("repository", "", "The repository that contains the package with the requested readme file.")
		codeartifact_getPackageVersionReadmeCmd.MarkFlagRequired("domain")
		codeartifact_getPackageVersionReadmeCmd.MarkFlagRequired("format")
		codeartifact_getPackageVersionReadmeCmd.MarkFlagRequired("package")
		codeartifact_getPackageVersionReadmeCmd.MarkFlagRequired("package-version")
		codeartifact_getPackageVersionReadmeCmd.MarkFlagRequired("repository")
	})
	codeartifactCmd.AddCommand(codeartifact_getPackageVersionReadmeCmd)
}
