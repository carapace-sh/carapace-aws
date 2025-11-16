package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_listPackageGroupsCmd = &cobra.Command{
	Use:   "list-package-groups",
	Short: "Returns a list of package groups in the requested domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_listPackageGroupsCmd).Standalone()

	codeartifact_listPackageGroupsCmd.Flags().String("domain", "", "The domain for which you want to list package groups.")
	codeartifact_listPackageGroupsCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
	codeartifact_listPackageGroupsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	codeartifact_listPackageGroupsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	codeartifact_listPackageGroupsCmd.Flags().String("prefix", "", "A prefix for which to search package groups.")
	codeartifact_listPackageGroupsCmd.MarkFlagRequired("domain")
	codeartifactCmd.AddCommand(codeartifact_listPackageGroupsCmd)
}
