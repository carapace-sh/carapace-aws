package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var grafana_listWorkspaceServiceAccountTokensCmd = &cobra.Command{
	Use:   "list-workspace-service-account-tokens",
	Short: "Returns a list of tokens for a workspace service account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grafana_listWorkspaceServiceAccountTokensCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(grafana_listWorkspaceServiceAccountTokensCmd).Standalone()

		grafana_listWorkspaceServiceAccountTokensCmd.Flags().String("max-results", "", "The maximum number of tokens to include in the results.")
		grafana_listWorkspaceServiceAccountTokensCmd.Flags().String("next-token", "", "The token for the next set of service accounts to return.")
		grafana_listWorkspaceServiceAccountTokensCmd.Flags().String("service-account-id", "", "The ID of the service account for which to return tokens.")
		grafana_listWorkspaceServiceAccountTokensCmd.Flags().String("workspace-id", "", "The ID of the workspace for which to return tokens.")
		grafana_listWorkspaceServiceAccountTokensCmd.MarkFlagRequired("service-account-id")
		grafana_listWorkspaceServiceAccountTokensCmd.MarkFlagRequired("workspace-id")
	})
	grafanaCmd.AddCommand(grafana_listWorkspaceServiceAccountTokensCmd)
}
