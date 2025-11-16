package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_listHitsCmd = &cobra.Command{
	Use:   "list-hits",
	Short: "The `ListHITs` operation returns all of a Requester's HITs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_listHitsCmd).Standalone()

	mturk_listHitsCmd.Flags().String("max-results", "", "")
	mturk_listHitsCmd.Flags().String("next-token", "", "Pagination token")
	mturkCmd.AddCommand(mturk_listHitsCmd)
}
