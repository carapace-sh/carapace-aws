package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcZonalShift_deletePracticeRunConfigurationCmd = &cobra.Command{
	Use:   "delete-practice-run-configuration",
	Short: "Deletes the practice run configuration for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcZonalShift_deletePracticeRunConfigurationCmd).Standalone()

	arcZonalShift_deletePracticeRunConfigurationCmd.Flags().String("resource-identifier", "", "The identifier for the resource that you want to delete the practice run configuration for.")
	arcZonalShift_deletePracticeRunConfigurationCmd.MarkFlagRequired("resource-identifier")
	arcZonalShiftCmd.AddCommand(arcZonalShift_deletePracticeRunConfigurationCmd)
}
