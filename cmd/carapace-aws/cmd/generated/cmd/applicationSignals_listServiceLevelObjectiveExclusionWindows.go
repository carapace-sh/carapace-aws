package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationSignals_listServiceLevelObjectiveExclusionWindowsCmd = &cobra.Command{
	Use:   "list-service-level-objective-exclusion-windows",
	Short: "Retrieves all exclusion windows configured for a specific SLO.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationSignals_listServiceLevelObjectiveExclusionWindowsCmd).Standalone()

	applicationSignals_listServiceLevelObjectiveExclusionWindowsCmd.Flags().String("id", "", "The ID of the SLO to list exclusion windows for.")
	applicationSignals_listServiceLevelObjectiveExclusionWindowsCmd.Flags().String("max-results", "", "The maximum number of results to return in one operation.")
	applicationSignals_listServiceLevelObjectiveExclusionWindowsCmd.Flags().String("next-token", "", "Include this value, if it was returned by the previous operation, to get the next set of service level objectives.")
	applicationSignals_listServiceLevelObjectiveExclusionWindowsCmd.MarkFlagRequired("id")
	applicationSignalsCmd.AddCommand(applicationSignals_listServiceLevelObjectiveExclusionWindowsCmd)
}
