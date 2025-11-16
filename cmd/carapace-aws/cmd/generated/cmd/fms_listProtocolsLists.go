package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_listProtocolsListsCmd = &cobra.Command{
	Use:   "list-protocols-lists",
	Short: "Returns an array of `ProtocolsListDataSummary` objects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_listProtocolsListsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fms_listProtocolsListsCmd).Standalone()

		fms_listProtocolsListsCmd.Flags().Bool("default-lists", false, "Specifies whether the lists to retrieve are default lists owned by Firewall Manager.")
		fms_listProtocolsListsCmd.Flags().String("max-results", "", "The maximum number of objects that you want Firewall Manager to return for this request.")
		fms_listProtocolsListsCmd.Flags().String("next-token", "", "If you specify a value for `MaxResults` in your list request, and you have more objects than the maximum, Firewall Manager returns this token in the response.")
		fms_listProtocolsListsCmd.Flags().Bool("no-default-lists", false, "Specifies whether the lists to retrieve are default lists owned by Firewall Manager.")
		fms_listProtocolsListsCmd.MarkFlagRequired("max-results")
		fms_listProtocolsListsCmd.Flag("no-default-lists").Hidden = true
	})
	fmsCmd.AddCommand(fms_listProtocolsListsCmd)
}
