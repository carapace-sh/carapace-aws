package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_deleteApplicationCmd = &cobra.Command{
	Use:   "delete-application",
	Short: "Delete application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_deleteApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_deleteApplicationCmd).Standalone()

		mgn_deleteApplicationCmd.Flags().String("account-id", "", "Account ID.")
		mgn_deleteApplicationCmd.Flags().String("application-id", "", "Application ID.")
		mgn_deleteApplicationCmd.MarkFlagRequired("application-id")
	})
	mgnCmd.AddCommand(mgn_deleteApplicationCmd)
}
