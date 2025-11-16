package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_deleteRunCmd = &cobra.Command{
	Use:   "delete-run",
	Short: "Deletes a run and returns a response with no body if the operation is successful.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_deleteRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_deleteRunCmd).Standalone()

		omics_deleteRunCmd.Flags().String("id", "", "The run's ID.")
		omics_deleteRunCmd.MarkFlagRequired("id")
	})
	omicsCmd.AddCommand(omics_deleteRunCmd)
}
