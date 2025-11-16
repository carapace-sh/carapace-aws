package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discovery_startBatchDeleteConfigurationTaskCmd = &cobra.Command{
	Use:   "start-batch-delete-configuration-task",
	Short: "Takes a list of configurationId as input and starts an asynchronous deletion task to remove the configurationItems.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discovery_startBatchDeleteConfigurationTaskCmd).Standalone()

	discovery_startBatchDeleteConfigurationTaskCmd.Flags().String("configuration-ids", "", "The list of configuration IDs that will be deleted by the task.")
	discovery_startBatchDeleteConfigurationTaskCmd.Flags().String("configuration-type", "", "The type of configuration item to delete.")
	discovery_startBatchDeleteConfigurationTaskCmd.MarkFlagRequired("configuration-ids")
	discovery_startBatchDeleteConfigurationTaskCmd.MarkFlagRequired("configuration-type")
	discoveryCmd.AddCommand(discovery_startBatchDeleteConfigurationTaskCmd)
}
