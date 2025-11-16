package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listPredefinedAttributesCmd = &cobra.Command{
	Use:   "list-predefined-attributes",
	Short: "Lists predefined attributes for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listPredefinedAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_listPredefinedAttributesCmd).Standalone()

		connect_listPredefinedAttributesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_listPredefinedAttributesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_listPredefinedAttributesCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_listPredefinedAttributesCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_listPredefinedAttributesCmd)
}
