package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mpa_cancelSessionCmd = &cobra.Command{
	Use:   "cancel-session",
	Short: "Cancels an approval session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mpa_cancelSessionCmd).Standalone()

	mpa_cancelSessionCmd.Flags().String("session-arn", "", "Amazon Resource Name (ARN) for the session.")
	mpa_cancelSessionCmd.MarkFlagRequired("session-arn")
	mpaCmd.AddCommand(mpa_cancelSessionCmd)
}
