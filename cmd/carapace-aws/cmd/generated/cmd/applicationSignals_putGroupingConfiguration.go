package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationSignals_putGroupingConfigurationCmd = &cobra.Command{
	Use:   "put-grouping-configuration",
	Short: "Creates or updates a grouping configuration that defines how services are organized and grouped in Application Signals dashboards and service maps.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationSignals_putGroupingConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationSignals_putGroupingConfigurationCmd).Standalone()

		applicationSignals_putGroupingConfigurationCmd.Flags().String("grouping-attribute-definitions", "", "An array of grouping attribute definitions that specify how services should be grouped.")
		applicationSignals_putGroupingConfigurationCmd.MarkFlagRequired("grouping-attribute-definitions")
	})
	applicationSignalsCmd.AddCommand(applicationSignals_putGroupingConfigurationCmd)
}
