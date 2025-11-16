package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_describeConnectionAliasPermissionsCmd = &cobra.Command{
	Use:   "describe-connection-alias-permissions",
	Short: "Describes the permissions that the owner of a connection alias has granted to another Amazon Web Services account for the specified connection alias.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_describeConnectionAliasPermissionsCmd).Standalone()

	workspaces_describeConnectionAliasPermissionsCmd.Flags().String("alias-id", "", "The identifier of the connection alias.")
	workspaces_describeConnectionAliasPermissionsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	workspaces_describeConnectionAliasPermissionsCmd.Flags().String("next-token", "", "If you received a `NextToken` from a previous call that was paginated, provide this token to receive the next set of results.")
	workspaces_describeConnectionAliasPermissionsCmd.MarkFlagRequired("alias-id")
	workspacesCmd.AddCommand(workspaces_describeConnectionAliasPermissionsCmd)
}
