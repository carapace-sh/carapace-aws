package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationSignals_deleteGroupingConfigurationCmd = &cobra.Command{
	Use:   "delete-grouping-configuration",
	Short: "Deletes a grouping configuration that defines how services are grouped and organized in Application Signals.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationSignals_deleteGroupingConfigurationCmd).Standalone()

	applicationSignalsCmd.AddCommand(applicationSignals_deleteGroupingConfigurationCmd)
}
