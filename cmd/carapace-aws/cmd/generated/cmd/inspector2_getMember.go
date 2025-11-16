package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_getMemberCmd = &cobra.Command{
	Use:   "get-member",
	Short: "Gets member information for your organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_getMemberCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_getMemberCmd).Standalone()

		inspector2_getMemberCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the member account to retrieve information on.")
		inspector2_getMemberCmd.MarkFlagRequired("account-id")
	})
	inspector2Cmd.AddCommand(inspector2_getMemberCmd)
}
