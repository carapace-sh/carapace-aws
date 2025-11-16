package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var grafana_listWorkspaceServiceAccountsCmd = &cobra.Command{
	Use:   "list-workspace-service-accounts",
	Short: "Returns a list of service accounts for a workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grafana_listWorkspaceServiceAccountsCmd).Standalone()

	grafana_listWorkspaceServiceAccountsCmd.Flags().String("max-results", "", "The maximum number of service accounts to include in the results.")
	grafana_listWorkspaceServiceAccountsCmd.Flags().String("next-token", "", "The token for the next set of service accounts to return.")
	grafana_listWorkspaceServiceAccountsCmd.Flags().String("workspace-id", "", "The workspace for which to list service accounts.")
	grafana_listWorkspaceServiceAccountsCmd.MarkFlagRequired("workspace-id")
	grafanaCmd.AddCommand(grafana_listWorkspaceServiceAccountsCmd)
}
