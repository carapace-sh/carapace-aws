package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_updateApplicationCmd = &cobra.Command{
	Use:   "update-application",
	Short: "Update application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_updateApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_updateApplicationCmd).Standalone()

		mgn_updateApplicationCmd.Flags().String("account-id", "", "Account ID.")
		mgn_updateApplicationCmd.Flags().String("application-id", "", "Application ID.")
		mgn_updateApplicationCmd.Flags().String("description", "", "Application description.")
		mgn_updateApplicationCmd.Flags().String("name", "", "Application name.")
		mgn_updateApplicationCmd.MarkFlagRequired("application-id")
	})
	mgnCmd.AddCommand(mgn_updateApplicationCmd)
}
