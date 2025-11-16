package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var grafana_describeWorkspaceAuthenticationCmd = &cobra.Command{
	Use:   "describe-workspace-authentication",
	Short: "Displays information about the authentication methods used in one Amazon Managed Grafana workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grafana_describeWorkspaceAuthenticationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(grafana_describeWorkspaceAuthenticationCmd).Standalone()

		grafana_describeWorkspaceAuthenticationCmd.Flags().String("workspace-id", "", "The ID of the workspace to return authentication information about.")
		grafana_describeWorkspaceAuthenticationCmd.MarkFlagRequired("workspace-id")
	})
	grafanaCmd.AddCommand(grafana_describeWorkspaceAuthenticationCmd)
}
