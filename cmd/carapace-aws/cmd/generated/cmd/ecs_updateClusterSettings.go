package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_updateClusterSettingsCmd = &cobra.Command{
	Use:   "update-cluster-settings",
	Short: "Modifies the settings to use for a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_updateClusterSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_updateClusterSettingsCmd).Standalone()

		ecs_updateClusterSettingsCmd.Flags().String("cluster", "", "The name of the cluster to modify the settings for.")
		ecs_updateClusterSettingsCmd.Flags().String("settings", "", "The setting to use by default for a cluster.")
		ecs_updateClusterSettingsCmd.MarkFlagRequired("cluster")
		ecs_updateClusterSettingsCmd.MarkFlagRequired("settings")
	})
	ecsCmd.AddCommand(ecs_updateClusterSettingsCmd)
}
