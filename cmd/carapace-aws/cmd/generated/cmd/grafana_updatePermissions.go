package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var grafana_updatePermissionsCmd = &cobra.Command{
	Use:   "update-permissions",
	Short: "Updates which users in a workspace have the Grafana `Admin` or `Editor` roles.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grafana_updatePermissionsCmd).Standalone()

	grafana_updatePermissionsCmd.Flags().String("update-instruction-batch", "", "An array of structures that contain the permission updates to make.")
	grafana_updatePermissionsCmd.Flags().String("workspace-id", "", "The ID of the workspace to update.")
	grafana_updatePermissionsCmd.MarkFlagRequired("update-instruction-batch")
	grafana_updatePermissionsCmd.MarkFlagRequired("workspace-id")
	grafanaCmd.AddCommand(grafana_updatePermissionsCmd)
}
