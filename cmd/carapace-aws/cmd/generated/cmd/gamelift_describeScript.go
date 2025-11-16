package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeScriptCmd = &cobra.Command{
	Use:   "describe-script",
	Short: "**This API works with the following fleet types:** EC2\n\nRetrieves properties for a Realtime script.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeScriptCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_describeScriptCmd).Standalone()

		gamelift_describeScriptCmd.Flags().String("script-id", "", "A unique identifier for the Realtime script to retrieve properties for.")
		gamelift_describeScriptCmd.MarkFlagRequired("script-id")
	})
	gameliftCmd.AddCommand(gamelift_describeScriptCmd)
}
