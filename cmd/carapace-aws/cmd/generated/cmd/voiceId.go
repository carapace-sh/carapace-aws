package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceIdCmd = &cobra.Command{
	Use:   "voice-id",
	Short: "Amazon Connect Voice ID provides real-time caller authentication and fraud risk detection, which make voice interactions in contact centers more secure and efficient.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceIdCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(voiceIdCmd).Standalone()

	})
	rootCmd.AddCommand(voiceIdCmd)
}
