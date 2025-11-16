package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var grafana_deleteWorkspaceCmd = &cobra.Command{
	Use:   "delete-workspace",
	Short: "Deletes an Amazon Managed Grafana workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grafana_deleteWorkspaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(grafana_deleteWorkspaceCmd).Standalone()

		grafana_deleteWorkspaceCmd.Flags().String("workspace-id", "", "The ID of the workspace to delete.")
		grafana_deleteWorkspaceCmd.MarkFlagRequired("workspace-id")
	})
	grafanaCmd.AddCommand(grafana_deleteWorkspaceCmd)
}
