package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discovery_describeBatchDeleteConfigurationTaskCmd = &cobra.Command{
	Use:   "describe-batch-delete-configuration-task",
	Short: "Takes a unique deletion task identifier as input and returns metadata about a configuration deletion task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discovery_describeBatchDeleteConfigurationTaskCmd).Standalone()

	discovery_describeBatchDeleteConfigurationTaskCmd.Flags().String("task-id", "", "The ID of the task to delete.")
	discovery_describeBatchDeleteConfigurationTaskCmd.MarkFlagRequired("task-id")
	discoveryCmd.AddCommand(discovery_describeBatchDeleteConfigurationTaskCmd)
}
