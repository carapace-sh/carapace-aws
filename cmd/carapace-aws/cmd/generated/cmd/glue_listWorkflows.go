package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_listWorkflowsCmd = &cobra.Command{
	Use:   "list-workflows",
	Short: "Lists names of workflows created in the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_listWorkflowsCmd).Standalone()

	glue_listWorkflowsCmd.Flags().String("max-results", "", "The maximum size of a list to return.")
	glue_listWorkflowsCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation request.")
	glueCmd.AddCommand(glue_listWorkflowsCmd)
}
