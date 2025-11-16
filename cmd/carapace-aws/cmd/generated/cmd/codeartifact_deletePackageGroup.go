package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_deletePackageGroupCmd = &cobra.Command{
	Use:   "delete-package-group",
	Short: "Deletes a package group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_deletePackageGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeartifact_deletePackageGroupCmd).Standalone()

		codeartifact_deletePackageGroupCmd.Flags().String("domain", "", "The domain that contains the package group to be deleted.")
		codeartifact_deletePackageGroupCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
		codeartifact_deletePackageGroupCmd.Flags().String("package-group", "", "The pattern of the package group to be deleted.")
		codeartifact_deletePackageGroupCmd.MarkFlagRequired("domain")
		codeartifact_deletePackageGroupCmd.MarkFlagRequired("package-group")
	})
	codeartifactCmd.AddCommand(codeartifact_deletePackageGroupCmd)
}
