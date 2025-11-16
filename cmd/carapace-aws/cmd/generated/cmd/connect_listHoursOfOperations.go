package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listHoursOfOperationsCmd = &cobra.Command{
	Use:   "list-hours-of-operations",
	Short: "Provides information about the hours of operation for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listHoursOfOperationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_listHoursOfOperationsCmd).Standalone()

		connect_listHoursOfOperationsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_listHoursOfOperationsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_listHoursOfOperationsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_listHoursOfOperationsCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_listHoursOfOperationsCmd)
}
