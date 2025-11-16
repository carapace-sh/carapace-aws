package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_updatePackageGroupCmd = &cobra.Command{
	Use:   "update-package-group",
	Short: "Updates a package group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_updatePackageGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeartifact_updatePackageGroupCmd).Standalone()

		codeartifact_updatePackageGroupCmd.Flags().String("contact-info", "", "Contact information which you want to update the requested package group with.")
		codeartifact_updatePackageGroupCmd.Flags().String("description", "", "The description you want to update the requested package group with.")
		codeartifact_updatePackageGroupCmd.Flags().String("domain", "", "The name of the domain which contains the package group to be updated.")
		codeartifact_updatePackageGroupCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
		codeartifact_updatePackageGroupCmd.Flags().String("package-group", "", "The pattern of the package group to be updated.")
		codeartifact_updatePackageGroupCmd.MarkFlagRequired("domain")
		codeartifact_updatePackageGroupCmd.MarkFlagRequired("package-group")
	})
	codeartifactCmd.AddCommand(codeartifact_updatePackageGroupCmd)
}
