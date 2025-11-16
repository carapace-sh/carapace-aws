package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearning_deleteMlmodelCmd = &cobra.Command{
	Use:   "delete-mlmodel",
	Short: "Assigns the `DELETED` status to an `MLModel`, rendering it unusable.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearning_deleteMlmodelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(machinelearning_deleteMlmodelCmd).Standalone()

		machinelearning_deleteMlmodelCmd.Flags().String("mlmodel-id", "", "A user-supplied ID that uniquely identifies the `MLModel`.")
		machinelearning_deleteMlmodelCmd.MarkFlagRequired("mlmodel-id")
	})
	machinelearningCmd.AddCommand(machinelearning_deleteMlmodelCmd)
}
