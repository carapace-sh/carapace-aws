package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_describePackageVersionCmd = &cobra.Command{
	Use:   "describe-package-version",
	Short: "Returns a [PackageVersionDescription](https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageVersionDescription.html) object that contains information about the requested package version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_describePackageVersionCmd).Standalone()

	codeartifact_describePackageVersionCmd.Flags().String("domain", "", "The name of the domain that contains the repository that contains the package version.")
	codeartifact_describePackageVersionCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
	codeartifact_describePackageVersionCmd.Flags().String("format", "", "A format that specifies the type of the requested package version.")
	codeartifact_describePackageVersionCmd.Flags().String("namespace", "", "The namespace of the requested package version.")
	codeartifact_describePackageVersionCmd.Flags().String("package", "", "The name of the requested package version.")
	codeartifact_describePackageVersionCmd.Flags().String("package-version", "", "A string that contains the package version (for example, `3.5.2`).")
	codeartifact_describePackageVersionCmd.Flags().String("repository", "", "The name of the repository that contains the package version.")
	codeartifact_describePackageVersionCmd.MarkFlagRequired("domain")
	codeartifact_describePackageVersionCmd.MarkFlagRequired("format")
	codeartifact_describePackageVersionCmd.MarkFlagRequired("package")
	codeartifact_describePackageVersionCmd.MarkFlagRequired("package-version")
	codeartifact_describePackageVersionCmd.MarkFlagRequired("repository")
	codeartifactCmd.AddCommand(codeartifact_describePackageVersionCmd)
}
