package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_getAssociatedPackageGroupCmd = &cobra.Command{
	Use:   "get-associated-package-group",
	Short: "Returns the most closely associated package group to the specified package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_getAssociatedPackageGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeartifact_getAssociatedPackageGroupCmd).Standalone()

		codeartifact_getAssociatedPackageGroupCmd.Flags().String("domain", "", "The name of the domain that contains the package from which to get the associated package group.")
		codeartifact_getAssociatedPackageGroupCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
		codeartifact_getAssociatedPackageGroupCmd.Flags().String("format", "", "The format of the package from which to get the associated package group.")
		codeartifact_getAssociatedPackageGroupCmd.Flags().String("namespace", "", "The namespace of the package from which to get the associated package group.")
		codeartifact_getAssociatedPackageGroupCmd.Flags().String("package", "", "The package from which to get the associated package group.")
		codeartifact_getAssociatedPackageGroupCmd.MarkFlagRequired("domain")
		codeartifact_getAssociatedPackageGroupCmd.MarkFlagRequired("format")
		codeartifact_getAssociatedPackageGroupCmd.MarkFlagRequired("package")
	})
	codeartifactCmd.AddCommand(codeartifact_getAssociatedPackageGroupCmd)
}
