package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteWorkflowCmd = &cobra.Command{
	Use:   "delete-workflow",
	Short: "Deletes a workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteWorkflowCmd).Standalone()

	glue_deleteWorkflowCmd.Flags().String("name", "", "Name of the workflow to be deleted.")
	glue_deleteWorkflowCmd.MarkFlagRequired("name")
	glueCmd.AddCommand(glue_deleteWorkflowCmd)
}
