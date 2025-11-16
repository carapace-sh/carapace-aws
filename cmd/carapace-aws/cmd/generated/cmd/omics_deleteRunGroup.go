package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_deleteRunGroupCmd = &cobra.Command{
	Use:   "delete-run-group",
	Short: "Deletes a run group and returns a response with no body if the operation is successful.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_deleteRunGroupCmd).Standalone()

	omics_deleteRunGroupCmd.Flags().String("id", "", "The run group's ID.")
	omics_deleteRunGroupCmd.MarkFlagRequired("id")
	omicsCmd.AddCommand(omics_deleteRunGroupCmd)
}
