package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_closeAccountCmd = &cobra.Command{
	Use:   "close-account",
	Short: "Closes an Amazon Web Services member account within an organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_closeAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(organizations_closeAccountCmd).Standalone()

		organizations_closeAccountCmd.Flags().String("account-id", "", "Retrieves the Amazon Web Services account Id for the current `CloseAccount` API request.")
		organizations_closeAccountCmd.MarkFlagRequired("account-id")
	})
	organizationsCmd.AddCommand(organizations_closeAccountCmd)
}
