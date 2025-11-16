package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_deletePackageCmd = &cobra.Command{
	Use:   "delete-package",
	Short: "Deletes a package and all associated package versions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_deletePackageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeartifact_deletePackageCmd).Standalone()

		codeartifact_deletePackageCmd.Flags().String("domain", "", "The name of the domain that contains the package to delete.")
		codeartifact_deletePackageCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
		codeartifact_deletePackageCmd.Flags().String("format", "", "The format of the requested package to delete.")
		codeartifact_deletePackageCmd.Flags().String("namespace", "", "The namespace of the package to delete.")
		codeartifact_deletePackageCmd.Flags().String("package", "", "The name of the package to delete.")
		codeartifact_deletePackageCmd.Flags().String("repository", "", "The name of the repository that contains the package to delete.")
		codeartifact_deletePackageCmd.MarkFlagRequired("domain")
		codeartifact_deletePackageCmd.MarkFlagRequired("format")
		codeartifact_deletePackageCmd.MarkFlagRequired("package")
		codeartifact_deletePackageCmd.MarkFlagRequired("repository")
	})
	codeartifactCmd.AddCommand(codeartifact_deletePackageCmd)
}
