package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fis_listActionsCmd = &cobra.Command{
	Use:   "list-actions",
	Short: "Lists the available FIS actions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fis_listActionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fis_listActionsCmd).Standalone()

		fis_listActionsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		fis_listActionsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	})
	fisCmd.AddCommand(fis_listActionsCmd)
}
