package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_searchEmailAddressesCmd = &cobra.Command{
	Use:   "search-email-addresses",
	Short: "Searches email address in an instance, with optional filtering.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_searchEmailAddressesCmd).Standalone()

	connect_searchEmailAddressesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_searchEmailAddressesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connect_searchEmailAddressesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_searchEmailAddressesCmd.Flags().String("search-criteria", "", "The search criteria to be used to return email addresses.")
	connect_searchEmailAddressesCmd.Flags().String("search-filter", "", "Filters to be applied to search results.")
	connect_searchEmailAddressesCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_searchEmailAddressesCmd)
}
