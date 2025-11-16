package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_describeCreateAccountStatusCmd = &cobra.Command{
	Use:   "describe-create-account-status",
	Short: "Retrieves the current status of an asynchronous request to create an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_describeCreateAccountStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(organizations_describeCreateAccountStatusCmd).Standalone()

		organizations_describeCreateAccountStatusCmd.Flags().String("create-account-request-id", "", "Specifies the `Id` value that uniquely identifies the `CreateAccount` request.")
		organizations_describeCreateAccountStatusCmd.MarkFlagRequired("create-account-request-id")
	})
	organizationsCmd.AddCommand(organizations_describeCreateAccountStatusCmd)
}
