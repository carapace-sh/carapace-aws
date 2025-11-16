package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_getMemberCmd = &cobra.Command{
	Use:   "get-member",
	Short: "Retrieves information about an account that's associated with an Amazon Macie administrator account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_getMemberCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_getMemberCmd).Standalone()

		macie2_getMemberCmd.Flags().String("id", "", "The unique identifier for the Amazon Macie resource that the request applies to.")
		macie2_getMemberCmd.MarkFlagRequired("id")
	})
	macie2Cmd.AddCommand(macie2_getMemberCmd)
}
