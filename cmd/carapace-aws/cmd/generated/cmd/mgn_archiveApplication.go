package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_archiveApplicationCmd = &cobra.Command{
	Use:   "archive-application",
	Short: "Archive application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_archiveApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_archiveApplicationCmd).Standalone()

		mgn_archiveApplicationCmd.Flags().String("account-id", "", "Account ID.")
		mgn_archiveApplicationCmd.Flags().String("application-id", "", "Application ID.")
		mgn_archiveApplicationCmd.MarkFlagRequired("application-id")
	})
	mgnCmd.AddCommand(mgn_archiveApplicationCmd)
}
