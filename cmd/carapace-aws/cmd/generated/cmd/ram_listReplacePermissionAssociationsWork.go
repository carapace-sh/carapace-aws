package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_listReplacePermissionAssociationsWorkCmd = &cobra.Command{
	Use:   "list-replace-permission-associations-work",
	Short: "Retrieves the current status of the asynchronous tasks performed by RAM when you perform the [ReplacePermissionAssociationsWork]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_listReplacePermissionAssociationsWorkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ram_listReplacePermissionAssociationsWorkCmd).Standalone()

		ram_listReplacePermissionAssociationsWorkCmd.Flags().String("max-results", "", "Specifies the total number of results that you want included on each page of the response.")
		ram_listReplacePermissionAssociationsWorkCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
		ram_listReplacePermissionAssociationsWorkCmd.Flags().String("status", "", "Specifies that you want to see only the details about requests with a status that matches this value.")
		ram_listReplacePermissionAssociationsWorkCmd.Flags().String("work-ids", "", "A list of IDs.")
	})
	ramCmd.AddCommand(ram_listReplacePermissionAssociationsWorkCmd)
}
