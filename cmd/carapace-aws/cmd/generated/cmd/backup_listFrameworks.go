package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_listFrameworksCmd = &cobra.Command{
	Use:   "list-frameworks",
	Short: "Returns a list of all frameworks for an Amazon Web Services account and Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_listFrameworksCmd).Standalone()

	backup_listFrameworksCmd.Flags().String("max-results", "", "The number of desired results from 1 to 1000. Optional.")
	backup_listFrameworksCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which can be used to return the next set of items in the list.")
	backupCmd.AddCommand(backup_listFrameworksCmd)
}
