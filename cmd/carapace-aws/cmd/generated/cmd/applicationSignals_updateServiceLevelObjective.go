package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationSignals_updateServiceLevelObjectiveCmd = &cobra.Command{
	Use:   "update-service-level-objective",
	Short: "Updates an existing service level objective (SLO).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationSignals_updateServiceLevelObjectiveCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationSignals_updateServiceLevelObjectiveCmd).Standalone()

		applicationSignals_updateServiceLevelObjectiveCmd.Flags().String("burn-rate-configurations", "", "Use this array to create *burn rates* for this SLO.")
		applicationSignals_updateServiceLevelObjectiveCmd.Flags().String("description", "", "An optional description for the SLO.")
		applicationSignals_updateServiceLevelObjectiveCmd.Flags().String("goal", "", "A structure that contains the attributes that determine the goal of the SLO.")
		applicationSignals_updateServiceLevelObjectiveCmd.Flags().String("id", "", "The Amazon Resource Name (ARN) or name of the service level objective that you want to update.")
		applicationSignals_updateServiceLevelObjectiveCmd.Flags().String("request-based-sli-config", "", "If this SLO is a request-based SLO, this structure defines the information about what performance metric this SLO will monitor.")
		applicationSignals_updateServiceLevelObjectiveCmd.Flags().String("sli-config", "", "If this SLO is a period-based SLO, this structure defines the information about what performance metric this SLO will monitor.")
		applicationSignals_updateServiceLevelObjectiveCmd.MarkFlagRequired("id")
	})
	applicationSignalsCmd.AddCommand(applicationSignals_updateServiceLevelObjectiveCmd)
}
