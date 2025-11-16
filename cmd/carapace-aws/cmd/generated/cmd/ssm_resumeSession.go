package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_resumeSessionCmd = &cobra.Command{
	Use:   "resume-session",
	Short: "Reconnects a session to a managed node after it has been disconnected.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_resumeSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_resumeSessionCmd).Standalone()

		ssm_resumeSessionCmd.Flags().String("session-id", "", "The ID of the disconnected session to resume.")
		ssm_resumeSessionCmd.MarkFlagRequired("session-id")
	})
	ssmCmd.AddCommand(ssm_resumeSessionCmd)
}
