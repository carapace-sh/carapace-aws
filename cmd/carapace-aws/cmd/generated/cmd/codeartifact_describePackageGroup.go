package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_describePackageGroupCmd = &cobra.Command{
	Use:   "describe-package-group",
	Short: "Returns a [PackageGroupDescription](https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageGroupDescription.html) object that contains information about the requested package group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_describePackageGroupCmd).Standalone()

	codeartifact_describePackageGroupCmd.Flags().String("domain", "", "The name of the domain that contains the package group.")
	codeartifact_describePackageGroupCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
	codeartifact_describePackageGroupCmd.Flags().String("package-group", "", "The pattern of the requested package group.")
	codeartifact_describePackageGroupCmd.MarkFlagRequired("domain")
	codeartifact_describePackageGroupCmd.MarkFlagRequired("package-group")
	codeartifactCmd.AddCommand(codeartifact_describePackageGroupCmd)
}
