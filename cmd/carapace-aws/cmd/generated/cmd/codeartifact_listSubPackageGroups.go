package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_listSubPackageGroupsCmd = &cobra.Command{
	Use:   "list-sub-package-groups",
	Short: "Returns a list of direct children of the specified package group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_listSubPackageGroupsCmd).Standalone()

	codeartifact_listSubPackageGroupsCmd.Flags().String("domain", "", "The name of the domain which contains the package group from which to list sub package groups.")
	codeartifact_listSubPackageGroupsCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
	codeartifact_listSubPackageGroupsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	codeartifact_listSubPackageGroupsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	codeartifact_listSubPackageGroupsCmd.Flags().String("package-group", "", "The pattern of the package group from which to list sub package groups.")
	codeartifact_listSubPackageGroupsCmd.MarkFlagRequired("domain")
	codeartifact_listSubPackageGroupsCmd.MarkFlagRequired("package-group")
	codeartifactCmd.AddCommand(codeartifact_listSubPackageGroupsCmd)
}
