package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_deleteWorkflowCmd = &cobra.Command{
	Use:   "delete-workflow",
	Short: "Deletes a workflow by specifying its ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_deleteWorkflowCmd).Standalone()

	omics_deleteWorkflowCmd.Flags().String("id", "", "The workflow's ID.")
	omics_deleteWorkflowCmd.MarkFlagRequired("id")
	omicsCmd.AddCommand(omics_deleteWorkflowCmd)
}
