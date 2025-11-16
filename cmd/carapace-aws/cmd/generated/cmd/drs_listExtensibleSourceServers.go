package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_listExtensibleSourceServersCmd = &cobra.Command{
	Use:   "list-extensible-source-servers",
	Short: "Returns a list of source servers on a staging account that are extensible, which means that: a.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_listExtensibleSourceServersCmd).Standalone()

	drs_listExtensibleSourceServersCmd.Flags().String("max-results", "", "The maximum number of extensible source servers to retrieve.")
	drs_listExtensibleSourceServersCmd.Flags().String("next-token", "", "The token of the next extensible source server to retrieve.")
	drs_listExtensibleSourceServersCmd.Flags().String("staging-account-id", "", "The Id of the staging Account to retrieve extensible source servers from.")
	drs_listExtensibleSourceServersCmd.MarkFlagRequired("staging-account-id")
	drsCmd.AddCommand(drs_listExtensibleSourceServersCmd)
}
