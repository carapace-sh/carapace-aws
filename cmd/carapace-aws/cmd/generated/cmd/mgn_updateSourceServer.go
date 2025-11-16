package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_updateSourceServerCmd = &cobra.Command{
	Use:   "update-source-server",
	Short: "Update Source Server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_updateSourceServerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_updateSourceServerCmd).Standalone()

		mgn_updateSourceServerCmd.Flags().String("account-id", "", "Update Source Server request account ID.")
		mgn_updateSourceServerCmd.Flags().String("connector-action", "", "Update Source Server request connector action.")
		mgn_updateSourceServerCmd.Flags().String("source-server-id", "", "Update Source Server request source server ID.")
		mgn_updateSourceServerCmd.MarkFlagRequired("source-server-id")
	})
	mgnCmd.AddCommand(mgn_updateSourceServerCmd)
}
