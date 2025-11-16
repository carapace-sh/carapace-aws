package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_describeSourceServersCmd = &cobra.Command{
	Use:   "describe-source-servers",
	Short: "Retrieves all SourceServers or multiple SourceServers by ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_describeSourceServersCmd).Standalone()

	mgn_describeSourceServersCmd.Flags().String("account-id", "", "Request to filter Source Servers list by Accoun ID.")
	mgn_describeSourceServersCmd.Flags().String("filters", "", "Request to filter Source Servers list.")
	mgn_describeSourceServersCmd.Flags().String("max-results", "", "Request to filter Source Servers list by maximum results.")
	mgn_describeSourceServersCmd.Flags().String("next-token", "", "Request to filter Source Servers list by next token.")
	mgnCmd.AddCommand(mgn_describeSourceServersCmd)
}
