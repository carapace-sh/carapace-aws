package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_unarchiveApplicationCmd = &cobra.Command{
	Use:   "unarchive-application",
	Short: "Unarchive application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_unarchiveApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_unarchiveApplicationCmd).Standalone()

		mgn_unarchiveApplicationCmd.Flags().String("account-id", "", "Account ID.")
		mgn_unarchiveApplicationCmd.Flags().String("application-id", "", "Application ID.")
		mgn_unarchiveApplicationCmd.MarkFlagRequired("application-id")
	})
	mgnCmd.AddCommand(mgn_unarchiveApplicationCmd)
}
