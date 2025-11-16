package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_associateSourceServersCmd = &cobra.Command{
	Use:   "associate-source-servers",
	Short: "Associate source servers to application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_associateSourceServersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_associateSourceServersCmd).Standalone()

		mgn_associateSourceServersCmd.Flags().String("account-id", "", "Account ID.")
		mgn_associateSourceServersCmd.Flags().String("application-id", "", "Application ID.")
		mgn_associateSourceServersCmd.Flags().String("source-server-ids", "", "Source server IDs list.")
		mgn_associateSourceServersCmd.MarkFlagRequired("application-id")
		mgn_associateSourceServersCmd.MarkFlagRequired("source-server-ids")
	})
	mgnCmd.AddCommand(mgn_associateSourceServersCmd)
}
