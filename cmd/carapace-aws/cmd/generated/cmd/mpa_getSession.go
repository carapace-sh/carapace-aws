package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mpa_getSessionCmd = &cobra.Command{
	Use:   "get-session",
	Short: "Returns details for an approval session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mpa_getSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mpa_getSessionCmd).Standalone()

		mpa_getSessionCmd.Flags().String("session-arn", "", "Amazon Resource Name (ARN) for the session.")
		mpa_getSessionCmd.MarkFlagRequired("session-arn")
	})
	mpaCmd.AddCommand(mpa_getSessionCmd)
}
