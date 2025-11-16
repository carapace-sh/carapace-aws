package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_disassociateApplicationsCmd = &cobra.Command{
	Use:   "disassociate-applications",
	Short: "Disassociate applications from wave.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_disassociateApplicationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_disassociateApplicationsCmd).Standalone()

		mgn_disassociateApplicationsCmd.Flags().String("account-id", "", "Account ID.")
		mgn_disassociateApplicationsCmd.Flags().String("application-ids", "", "Application IDs list.")
		mgn_disassociateApplicationsCmd.Flags().String("wave-id", "", "Wave ID.")
		mgn_disassociateApplicationsCmd.MarkFlagRequired("application-ids")
		mgn_disassociateApplicationsCmd.MarkFlagRequired("wave-id")
	})
	mgnCmd.AddCommand(mgn_disassociateApplicationsCmd)
}
