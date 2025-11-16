package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationSignals_createServiceLevelObjectiveCmd = &cobra.Command{
	Use:   "create-service-level-objective",
	Short: "Creates a service level objective (SLO), which can help you ensure that your critical business operations are meeting customer expectations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationSignals_createServiceLevelObjectiveCmd).Standalone()

	applicationSignals_createServiceLevelObjectiveCmd.Flags().String("burn-rate-configurations", "", "Use this array to create *burn rates* for this SLO.")
	applicationSignals_createServiceLevelObjectiveCmd.Flags().String("description", "", "An optional description for this SLO.")
	applicationSignals_createServiceLevelObjectiveCmd.Flags().String("goal", "", "This structure contains the attributes that determine the goal of the SLO.")
	applicationSignals_createServiceLevelObjectiveCmd.Flags().String("name", "", "A name for this SLO.")
	applicationSignals_createServiceLevelObjectiveCmd.Flags().String("request-based-sli-config", "", "If this SLO is a request-based SLO, this structure defines the information about what performance metric this SLO will monitor.")
	applicationSignals_createServiceLevelObjectiveCmd.Flags().String("sli-config", "", "If this SLO is a period-based SLO, this structure defines the information about what performance metric this SLO will monitor.")
	applicationSignals_createServiceLevelObjectiveCmd.Flags().String("tags", "", "A list of key-value pairs to associate with the SLO.")
	applicationSignals_createServiceLevelObjectiveCmd.MarkFlagRequired("name")
	applicationSignalsCmd.AddCommand(applicationSignals_createServiceLevelObjectiveCmd)
}
