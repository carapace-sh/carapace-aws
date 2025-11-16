package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_updateEventSourcesConfigCmd = &cobra.Command{
	Use:   "update-event-sources-config",
	Short: "Enables or disables integration with a service that can be integrated with DevOps Guru.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_updateEventSourcesConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devopsGuru_updateEventSourcesConfigCmd).Standalone()

		devopsGuru_updateEventSourcesConfigCmd.Flags().String("event-sources", "", "Configuration information about the integration of DevOps Guru as the Consumer via EventBridge with another AWS Service.")
	})
	devopsGuruCmd.AddCommand(devopsGuru_updateEventSourcesConfigCmd)
}
