package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_getArtifactUrlCmd = &cobra.Command{
	Use:   "get-artifact-url",
	Short: "Returns the artifact info that corresponds to an artifact id.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_getArtifactUrlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplify_getArtifactUrlCmd).Standalone()

		amplify_getArtifactUrlCmd.Flags().String("artifact-id", "", "The unique ID for an artifact.")
		amplify_getArtifactUrlCmd.MarkFlagRequired("artifact-id")
	})
	amplifyCmd.AddCommand(amplify_getArtifactUrlCmd)
}
