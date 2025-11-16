package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_listPlatformApplicationsCmd = &cobra.Command{
	Use:   "list-platform-applications",
	Short: "Lists the platform application objects for the supported push notification services, such as APNS and GCM (Firebase Cloud Messaging).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_listPlatformApplicationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sns_listPlatformApplicationsCmd).Standalone()

		sns_listPlatformApplicationsCmd.Flags().String("next-token", "", "`NextToken` string is used when calling `ListPlatformApplications` action to retrieve additional records that are available after the first page results.")
	})
	snsCmd.AddCommand(sns_listPlatformApplicationsCmd)
}
