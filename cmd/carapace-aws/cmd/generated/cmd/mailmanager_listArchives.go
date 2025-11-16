package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_listArchivesCmd = &cobra.Command{
	Use:   "list-archives",
	Short: "Returns a list of all email archives in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_listArchivesCmd).Standalone()

	mailmanager_listArchivesCmd.Flags().String("next-token", "", "If NextToken is returned, there are more results available.")
	mailmanager_listArchivesCmd.Flags().String("page-size", "", "The maximum number of archives that are returned per call.")
	mailmanagerCmd.AddCommand(mailmanager_listArchivesCmd)
}
