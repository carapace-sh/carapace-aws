package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_getWorkflowCmd = &cobra.Command{
	Use:   "get-workflow",
	Short: "Returns information about a workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_getWorkflowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecatalyst_getWorkflowCmd).Standalone()

		codecatalyst_getWorkflowCmd.Flags().String("id", "", "The ID of the workflow.")
		codecatalyst_getWorkflowCmd.Flags().String("project-name", "", "The name of the project in the space.")
		codecatalyst_getWorkflowCmd.Flags().String("space-name", "", "The name of the space.")
		codecatalyst_getWorkflowCmd.MarkFlagRequired("id")
		codecatalyst_getWorkflowCmd.MarkFlagRequired("project-name")
		codecatalyst_getWorkflowCmd.MarkFlagRequired("space-name")
	})
	codecatalystCmd.AddCommand(codecatalyst_getWorkflowCmd)
}
