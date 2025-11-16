package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_listAssociatedPackagesCmd = &cobra.Command{
	Use:   "list-associated-packages",
	Short: "Returns a list of packages associated with the requested package group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_listAssociatedPackagesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeartifact_listAssociatedPackagesCmd).Standalone()

		codeartifact_listAssociatedPackagesCmd.Flags().String("domain", "", "The name of the domain that contains the package group from which to list associated packages.")
		codeartifact_listAssociatedPackagesCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
		codeartifact_listAssociatedPackagesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		codeartifact_listAssociatedPackagesCmd.Flags().String("next-token", "", "The token for the next set of results.")
		codeartifact_listAssociatedPackagesCmd.Flags().String("package-group", "", "The pattern of the package group from which to list associated packages.")
		codeartifact_listAssociatedPackagesCmd.Flags().String("preview", "", "When this flag is included, `ListAssociatedPackages` will return a list of packages that would be associated with a package group, even if it does not exist.")
		codeartifact_listAssociatedPackagesCmd.MarkFlagRequired("domain")
		codeartifact_listAssociatedPackagesCmd.MarkFlagRequired("package-group")
	})
	codeartifactCmd.AddCommand(codeartifact_listAssociatedPackagesCmd)
}
