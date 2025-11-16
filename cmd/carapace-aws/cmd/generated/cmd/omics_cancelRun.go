package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_cancelRunCmd = &cobra.Command{
	Use:   "cancel-run",
	Short: "Cancels a run using its ID and returns a response with no body if the operation is successful.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_cancelRunCmd).Standalone()

	omics_cancelRunCmd.Flags().String("id", "", "The run's ID.")
	omics_cancelRunCmd.MarkFlagRequired("id")
	omicsCmd.AddCommand(omics_cancelRunCmd)
}
