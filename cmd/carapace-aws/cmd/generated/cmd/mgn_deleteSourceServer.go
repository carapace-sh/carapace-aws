package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_deleteSourceServerCmd = &cobra.Command{
	Use:   "delete-source-server",
	Short: "Deletes a single source server by ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_deleteSourceServerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_deleteSourceServerCmd).Standalone()

		mgn_deleteSourceServerCmd.Flags().String("account-id", "", "Request to delete Source Server from service by Account ID.")
		mgn_deleteSourceServerCmd.Flags().String("source-server-id", "", "Request to delete Source Server from service by Server ID.")
		mgn_deleteSourceServerCmd.MarkFlagRequired("source-server-id")
	})
	mgnCmd.AddCommand(mgn_deleteSourceServerCmd)
}
