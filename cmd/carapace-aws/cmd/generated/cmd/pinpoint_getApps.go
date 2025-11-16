package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getAppsCmd = &cobra.Command{
	Use:   "get-apps",
	Short: "Retrieves information about all the applications that are associated with your Amazon Pinpoint account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getAppsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_getAppsCmd).Standalone()

		pinpoint_getAppsCmd.Flags().String("page-size", "", "The maximum number of items to include in each page of a paginated response.")
		pinpoint_getAppsCmd.Flags().String("token", "", "The NextToken string that specifies which page of results to return in a paginated response.")
	})
	pinpointCmd.AddCommand(pinpoint_getAppsCmd)
}
