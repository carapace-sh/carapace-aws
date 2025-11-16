package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_describeConnectionAliasesCmd = &cobra.Command{
	Use:   "describe-connection-aliases",
	Short: "Retrieves a list that describes the connection aliases used for cross-Region redirection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_describeConnectionAliasesCmd).Standalone()

	workspaces_describeConnectionAliasesCmd.Flags().String("alias-ids", "", "The identifiers of the connection aliases to describe.")
	workspaces_describeConnectionAliasesCmd.Flags().String("limit", "", "The maximum number of connection aliases to return.")
	workspaces_describeConnectionAliasesCmd.Flags().String("next-token", "", "If you received a `NextToken` from a previous call that was paginated, provide this token to receive the next set of results.")
	workspaces_describeConnectionAliasesCmd.Flags().String("resource-id", "", "The identifier of the directory associated with the connection alias.")
	workspacesCmd.AddCommand(workspaces_describeConnectionAliasesCmd)
}
