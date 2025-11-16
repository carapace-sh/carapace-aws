package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_describeAccountCmd = &cobra.Command{
	Use:   "describe-account",
	Short: "Retrieves Organizations-related information about the specified account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_describeAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(organizations_describeAccountCmd).Standalone()

		organizations_describeAccountCmd.Flags().String("account-id", "", "The unique identifier (ID) of the Amazon Web Services account that you want information about.")
		organizations_describeAccountCmd.MarkFlagRequired("account-id")
	})
	organizationsCmd.AddCommand(organizations_describeAccountCmd)
}
