package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_disassociateSourceServersCmd = &cobra.Command{
	Use:   "disassociate-source-servers",
	Short: "Disassociate source servers from application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_disassociateSourceServersCmd).Standalone()

	mgn_disassociateSourceServersCmd.Flags().String("account-id", "", "Account ID.")
	mgn_disassociateSourceServersCmd.Flags().String("application-id", "", "Application ID.")
	mgn_disassociateSourceServersCmd.Flags().String("source-server-ids", "", "Source server IDs list.")
	mgn_disassociateSourceServersCmd.MarkFlagRequired("application-id")
	mgn_disassociateSourceServersCmd.MarkFlagRequired("source-server-ids")
	mgnCmd.AddCommand(mgn_disassociateSourceServersCmd)
}
