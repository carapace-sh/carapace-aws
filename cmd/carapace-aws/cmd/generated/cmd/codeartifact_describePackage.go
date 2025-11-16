package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_describePackageCmd = &cobra.Command{
	Use:   "describe-package",
	Short: "Returns a [PackageDescription](https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageDescription.html) object that contains information about the requested package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_describePackageCmd).Standalone()

	codeartifact_describePackageCmd.Flags().String("domain", "", "The name of the domain that contains the repository that contains the package.")
	codeartifact_describePackageCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
	codeartifact_describePackageCmd.Flags().String("format", "", "A format that specifies the type of the requested package.")
	codeartifact_describePackageCmd.Flags().String("namespace", "", "The namespace of the requested package.")
	codeartifact_describePackageCmd.Flags().String("package", "", "The name of the requested package.")
	codeartifact_describePackageCmd.Flags().String("repository", "", "The name of the repository that contains the requested package.")
	codeartifact_describePackageCmd.MarkFlagRequired("domain")
	codeartifact_describePackageCmd.MarkFlagRequired("format")
	codeartifact_describePackageCmd.MarkFlagRequired("package")
	codeartifact_describePackageCmd.MarkFlagRequired("repository")
	codeartifactCmd.AddCommand(codeartifact_describePackageCmd)
}
