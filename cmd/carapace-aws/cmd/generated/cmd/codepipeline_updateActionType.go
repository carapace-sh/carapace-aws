package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_updateActionTypeCmd = &cobra.Command{
	Use:   "update-action-type",
	Short: "Updates an action type that was created with any supported integration model, where the action type is to be used by customers of the action type provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_updateActionTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codepipeline_updateActionTypeCmd).Standalone()

		codepipeline_updateActionTypeCmd.Flags().String("action-type", "", "The action type definition for the action type to be updated.")
		codepipeline_updateActionTypeCmd.MarkFlagRequired("action-type")
	})
	codepipelineCmd.AddCommand(codepipeline_updateActionTypeCmd)
}
