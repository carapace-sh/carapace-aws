package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeconnections_getSyncConfigurationCmd = &cobra.Command{
	Use:   "get-sync-configuration",
	Short: "Returns details about a sync configuration, including the sync type and resource name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeconnections_getSyncConfigurationCmd).Standalone()

	codeconnections_getSyncConfigurationCmd.Flags().String("resource-name", "", "The name of the Amazon Web Services resource for the sync configuration for which you want to retrieve information.")
	codeconnections_getSyncConfigurationCmd.Flags().String("sync-type", "", "The sync type for the sync configuration for which you want to retrieve information.")
	codeconnections_getSyncConfigurationCmd.MarkFlagRequired("resource-name")
	codeconnections_getSyncConfigurationCmd.MarkFlagRequired("sync-type")
	codeconnectionsCmd.AddCommand(codeconnections_getSyncConfigurationCmd)
}
