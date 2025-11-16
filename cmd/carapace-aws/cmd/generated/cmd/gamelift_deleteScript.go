package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_deleteScriptCmd = &cobra.Command{
	Use:   "delete-script",
	Short: "**This API works with the following fleet types:** EC2\n\nDeletes a Realtime script.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_deleteScriptCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_deleteScriptCmd).Standalone()

		gamelift_deleteScriptCmd.Flags().String("script-id", "", "A unique identifier for the Realtime script to delete.")
		gamelift_deleteScriptCmd.MarkFlagRequired("script-id")
	})
	gameliftCmd.AddCommand(gamelift_deleteScriptCmd)
}
