package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_listAppsListsCmd = &cobra.Command{
	Use:   "list-apps-lists",
	Short: "Returns an array of `AppsListDataSummary` objects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_listAppsListsCmd).Standalone()

	fms_listAppsListsCmd.Flags().Bool("default-lists", false, "Specifies whether the lists to retrieve are default lists owned by Firewall Manager.")
	fms_listAppsListsCmd.Flags().String("max-results", "", "The maximum number of objects that you want Firewall Manager to return for this request.")
	fms_listAppsListsCmd.Flags().String("next-token", "", "If you specify a value for `MaxResults` in your list request, and you have more objects than the maximum, Firewall Manager returns this token in the response.")
	fms_listAppsListsCmd.Flags().Bool("no-default-lists", false, "Specifies whether the lists to retrieve are default lists owned by Firewall Manager.")
	fms_listAppsListsCmd.MarkFlagRequired("max-results")
	fms_listAppsListsCmd.Flag("no-default-lists").Hidden = true
	fmsCmd.AddCommand(fms_listAppsListsCmd)
}
