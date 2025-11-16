package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotData_listRetainedMessagesCmd = &cobra.Command{
	Use:   "list-retained-messages",
	Short: "Lists summary information about the retained messages stored for the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotData_listRetainedMessagesCmd).Standalone()

	iotData_listRetainedMessagesCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
	iotData_listRetainedMessagesCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
	iotDataCmd.AddCommand(iotData_listRetainedMessagesCmd)
}
