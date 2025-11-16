package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listHoursOfOperationOverridesCmd = &cobra.Command{
	Use:   "list-hours-of-operation-overrides",
	Short: "List the hours of operation overrides.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listHoursOfOperationOverridesCmd).Standalone()

	connect_listHoursOfOperationOverridesCmd.Flags().String("hours-of-operation-id", "", "The identifier for the hours of operation.")
	connect_listHoursOfOperationOverridesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_listHoursOfOperationOverridesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connect_listHoursOfOperationOverridesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_listHoursOfOperationOverridesCmd.MarkFlagRequired("hours-of-operation-id")
	connect_listHoursOfOperationOverridesCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_listHoursOfOperationOverridesCmd)
}
