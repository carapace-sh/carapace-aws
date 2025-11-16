package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var grafana_describeWorkspaceCmd = &cobra.Command{
	Use:   "describe-workspace",
	Short: "Displays information about one Amazon Managed Grafana workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grafana_describeWorkspaceCmd).Standalone()

	grafana_describeWorkspaceCmd.Flags().String("workspace-id", "", "The ID of the workspace to display information about.")
	grafana_describeWorkspaceCmd.MarkFlagRequired("workspace-id")
	grafanaCmd.AddCommand(grafana_describeWorkspaceCmd)
}
