package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listFarmsCmd = &cobra.Command{
	Use:   "list-farms",
	Short: "Lists farms.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listFarmsCmd).Standalone()

	deadline_listFarmsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	deadline_listFarmsCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
	deadline_listFarmsCmd.Flags().String("principal-id", "", "The principal ID of the member to list on the farm.")
	deadlineCmd.AddCommand(deadline_listFarmsCmd)
}
