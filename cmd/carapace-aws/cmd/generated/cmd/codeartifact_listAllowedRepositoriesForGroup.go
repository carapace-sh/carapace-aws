package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_listAllowedRepositoriesForGroupCmd = &cobra.Command{
	Use:   "list-allowed-repositories-for-group",
	Short: "Lists the repositories in the added repositories list of the specified restriction type for a package group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_listAllowedRepositoriesForGroupCmd).Standalone()

	codeartifact_listAllowedRepositoriesForGroupCmd.Flags().String("domain", "", "The name of the domain that contains the package group from which to list allowed repositories.")
	codeartifact_listAllowedRepositoriesForGroupCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
	codeartifact_listAllowedRepositoriesForGroupCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	codeartifact_listAllowedRepositoriesForGroupCmd.Flags().String("next-token", "", "The token for the next set of results.")
	codeartifact_listAllowedRepositoriesForGroupCmd.Flags().String("origin-restriction-type", "", "The origin configuration restriction type of which to list allowed repositories.")
	codeartifact_listAllowedRepositoriesForGroupCmd.Flags().String("package-group", "", "The pattern of the package group from which to list allowed repositories.")
	codeartifact_listAllowedRepositoriesForGroupCmd.MarkFlagRequired("domain")
	codeartifact_listAllowedRepositoriesForGroupCmd.MarkFlagRequired("origin-restriction-type")
	codeartifact_listAllowedRepositoriesForGroupCmd.MarkFlagRequired("package-group")
	codeartifactCmd.AddCommand(codeartifact_listAllowedRepositoriesForGroupCmd)
}
