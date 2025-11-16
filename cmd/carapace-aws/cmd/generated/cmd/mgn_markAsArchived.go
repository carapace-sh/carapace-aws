package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_markAsArchivedCmd = &cobra.Command{
	Use:   "mark-as-archived",
	Short: "Archives specific Source Servers by setting the SourceServer.isArchived property to true for specified SourceServers by ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_markAsArchivedCmd).Standalone()

	mgn_markAsArchivedCmd.Flags().String("account-id", "", "Mark as archived by Account ID.")
	mgn_markAsArchivedCmd.Flags().String("source-server-id", "", "Mark as archived by Source Server ID.")
	mgn_markAsArchivedCmd.MarkFlagRequired("source-server-id")
	mgnCmd.AddCommand(mgn_markAsArchivedCmd)
}
