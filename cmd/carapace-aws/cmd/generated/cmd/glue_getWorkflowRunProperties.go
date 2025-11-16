package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getWorkflowRunPropertiesCmd = &cobra.Command{
	Use:   "get-workflow-run-properties",
	Short: "Retrieves the workflow run properties which were set during the run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getWorkflowRunPropertiesCmd).Standalone()

	glue_getWorkflowRunPropertiesCmd.Flags().String("name", "", "Name of the workflow which was run.")
	glue_getWorkflowRunPropertiesCmd.Flags().String("run-id", "", "The ID of the workflow run whose run properties should be returned.")
	glue_getWorkflowRunPropertiesCmd.MarkFlagRequired("name")
	glue_getWorkflowRunPropertiesCmd.MarkFlagRequired("run-id")
	glueCmd.AddCommand(glue_getWorkflowRunPropertiesCmd)
}
