package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationSignals_deleteServiceLevelObjectiveCmd = &cobra.Command{
	Use:   "delete-service-level-objective",
	Short: "Deletes the specified service level objective.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationSignals_deleteServiceLevelObjectiveCmd).Standalone()

	applicationSignals_deleteServiceLevelObjectiveCmd.Flags().String("id", "", "The ARN or name of the service level objective to delete.")
	applicationSignals_deleteServiceLevelObjectiveCmd.MarkFlagRequired("id")
	applicationSignalsCmd.AddCommand(applicationSignals_deleteServiceLevelObjectiveCmd)
}
