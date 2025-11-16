package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_moveAccountCmd = &cobra.Command{
	Use:   "move-account",
	Short: "Moves an account from its current source parent root or organizational unit (OU) to the specified destination parent root or OU.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_moveAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(organizations_moveAccountCmd).Standalone()

		organizations_moveAccountCmd.Flags().String("account-id", "", "The unique identifier (ID) of the account that you want to move.")
		organizations_moveAccountCmd.Flags().String("destination-parent-id", "", "The unique identifier (ID) of the root or organizational unit that you want to move the account to.")
		organizations_moveAccountCmd.Flags().String("source-parent-id", "", "The unique identifier (ID) of the root or organizational unit that you want to move the account from.")
		organizations_moveAccountCmd.MarkFlagRequired("account-id")
		organizations_moveAccountCmd.MarkFlagRequired("destination-parent-id")
		organizations_moveAccountCmd.MarkFlagRequired("source-parent-id")
	})
	organizationsCmd.AddCommand(organizations_moveAccountCmd)
}
