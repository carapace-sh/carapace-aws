package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationSignals_getServiceLevelObjectiveCmd = &cobra.Command{
	Use:   "get-service-level-objective",
	Short: "Returns information about one SLO created in the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationSignals_getServiceLevelObjectiveCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationSignals_getServiceLevelObjectiveCmd).Standalone()

		applicationSignals_getServiceLevelObjectiveCmd.Flags().String("id", "", "The ARN or name of the SLO that you want to retrieve information about.")
		applicationSignals_getServiceLevelObjectiveCmd.MarkFlagRequired("id")
	})
	applicationSignalsCmd.AddCommand(applicationSignals_getServiceLevelObjectiveCmd)
}
