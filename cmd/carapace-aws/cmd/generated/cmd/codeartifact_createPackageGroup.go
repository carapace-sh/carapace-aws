package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_createPackageGroupCmd = &cobra.Command{
	Use:   "create-package-group",
	Short: "Creates a package group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_createPackageGroupCmd).Standalone()

	codeartifact_createPackageGroupCmd.Flags().String("contact-info", "", "The contact information for the created package group.")
	codeartifact_createPackageGroupCmd.Flags().String("description", "", "A description of the package group.")
	codeartifact_createPackageGroupCmd.Flags().String("domain", "", "The name of the domain in which you want to create a package group.")
	codeartifact_createPackageGroupCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
	codeartifact_createPackageGroupCmd.Flags().String("package-group", "", "The pattern of the package group to create.")
	codeartifact_createPackageGroupCmd.Flags().String("tags", "", "One or more tag key-value pairs for the package group.")
	codeartifact_createPackageGroupCmd.MarkFlagRequired("domain")
	codeartifact_createPackageGroupCmd.MarkFlagRequired("package-group")
	codeartifactCmd.AddCommand(codeartifact_createPackageGroupCmd)
}
