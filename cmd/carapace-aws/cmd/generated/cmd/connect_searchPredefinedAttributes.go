package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_searchPredefinedAttributesCmd = &cobra.Command{
	Use:   "search-predefined-attributes",
	Short: "Searches predefined attributes that meet certain criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_searchPredefinedAttributesCmd).Standalone()

	connect_searchPredefinedAttributesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_searchPredefinedAttributesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connect_searchPredefinedAttributesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_searchPredefinedAttributesCmd.Flags().String("search-criteria", "", "The search criteria to be used to return predefined attributes.")
	connect_searchPredefinedAttributesCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_searchPredefinedAttributesCmd)
}
