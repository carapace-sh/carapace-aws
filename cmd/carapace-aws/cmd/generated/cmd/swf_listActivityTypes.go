package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_listActivityTypesCmd = &cobra.Command{
	Use:   "list-activity-types",
	Short: "Returns information about all activities registered in the specified domain that match the specified name and registration status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_listActivityTypesCmd).Standalone()

	swf_listActivityTypesCmd.Flags().String("domain", "", "The name of the domain in which the activity types have been registered.")
	swf_listActivityTypesCmd.Flags().String("maximum-page-size", "", "The maximum number of results that are returned per call.")
	swf_listActivityTypesCmd.Flags().String("name", "", "If specified, only lists the activity types that have this name.")
	swf_listActivityTypesCmd.Flags().String("next-page-token", "", "If `NextPageToken` is returned there are more results available.")
	swf_listActivityTypesCmd.Flags().String("registration-status", "", "Specifies the registration status of the activity types to list.")
	swf_listActivityTypesCmd.Flags().String("reverse-order", "", "When set to `true`, returns the results in reverse order.")
	swf_listActivityTypesCmd.MarkFlagRequired("domain")
	swf_listActivityTypesCmd.MarkFlagRequired("registration-status")
	swfCmd.AddCommand(swf_listActivityTypesCmd)
}
