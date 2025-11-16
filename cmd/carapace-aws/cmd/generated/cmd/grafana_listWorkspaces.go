package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var grafana_listWorkspacesCmd = &cobra.Command{
	Use:   "list-workspaces",
	Short: "Returns a list of Amazon Managed Grafana workspaces in the account, with some information about each workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grafana_listWorkspacesCmd).Standalone()

	grafana_listWorkspacesCmd.Flags().String("max-results", "", "The maximum number of workspaces to include in the results.")
	grafana_listWorkspacesCmd.Flags().String("next-token", "", "The token for the next set of workspaces to return.")
	grafanaCmd.AddCommand(grafana_listWorkspacesCmd)
}
