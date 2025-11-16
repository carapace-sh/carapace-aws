package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_getCurrentUserDataCmd = &cobra.Command{
	Use:   "get-current-user-data",
	Short: "Gets the real-time active user data from the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_getCurrentUserDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_getCurrentUserDataCmd).Standalone()

		connect_getCurrentUserDataCmd.Flags().String("filters", "", "The filters to apply to returned user data.")
		connect_getCurrentUserDataCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_getCurrentUserDataCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_getCurrentUserDataCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_getCurrentUserDataCmd.MarkFlagRequired("filters")
		connect_getCurrentUserDataCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_getCurrentUserDataCmd)
}
