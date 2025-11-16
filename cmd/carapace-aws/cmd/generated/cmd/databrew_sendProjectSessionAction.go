package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_sendProjectSessionActionCmd = &cobra.Command{
	Use:   "send-project-session-action",
	Short: "Performs a recipe step within an interactive DataBrew session that's currently open.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_sendProjectSessionActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(databrew_sendProjectSessionActionCmd).Standalone()

		databrew_sendProjectSessionActionCmd.Flags().String("client-session-id", "", "A unique identifier for an interactive session that's currently open and ready for work.")
		databrew_sendProjectSessionActionCmd.Flags().String("name", "", "The name of the project to apply the action to.")
		databrew_sendProjectSessionActionCmd.Flags().String("preview", "", "If true, the result of the recipe step will be returned, but not applied.")
		databrew_sendProjectSessionActionCmd.Flags().String("recipe-step", "", "")
		databrew_sendProjectSessionActionCmd.Flags().String("step-index", "", "The index from which to preview a step.")
		databrew_sendProjectSessionActionCmd.Flags().String("view-frame", "", "")
		databrew_sendProjectSessionActionCmd.MarkFlagRequired("name")
	})
	databrewCmd.AddCommand(databrew_sendProjectSessionActionCmd)
}
