package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_updateMacieSessionCmd = &cobra.Command{
	Use:   "update-macie-session",
	Short: "Suspends or re-enables Amazon Macie, or updates the configuration settings for a Macie account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_updateMacieSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_updateMacieSessionCmd).Standalone()

		macie2_updateMacieSessionCmd.Flags().String("finding-publishing-frequency", "", "Specifies how often to publish updates to policy findings for the account.")
		macie2_updateMacieSessionCmd.Flags().String("status", "", "Specifies a new status for the account.")
	})
	macie2Cmd.AddCommand(macie2_updateMacieSessionCmd)
}
