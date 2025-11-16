package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_associateApplicationsCmd = &cobra.Command{
	Use:   "associate-applications",
	Short: "Associate applications to wave.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_associateApplicationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_associateApplicationsCmd).Standalone()

		mgn_associateApplicationsCmd.Flags().String("account-id", "", "Account ID.")
		mgn_associateApplicationsCmd.Flags().String("application-ids", "", "Application IDs list.")
		mgn_associateApplicationsCmd.Flags().String("wave-id", "", "Wave ID.")
		mgn_associateApplicationsCmd.MarkFlagRequired("application-ids")
		mgn_associateApplicationsCmd.MarkFlagRequired("wave-id")
	})
	mgnCmd.AddCommand(mgn_associateApplicationsCmd)
}
