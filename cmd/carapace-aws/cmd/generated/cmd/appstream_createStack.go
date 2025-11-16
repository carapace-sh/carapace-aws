package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_createStackCmd = &cobra.Command{
	Use:   "create-stack",
	Short: "Creates a stack to start streaming applications to users.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_createStackCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_createStackCmd).Standalone()

		appstream_createStackCmd.Flags().String("access-endpoints", "", "The list of interface VPC endpoint (interface endpoint) objects.")
		appstream_createStackCmd.Flags().String("application-settings", "", "The persistent application settings for users of a stack.")
		appstream_createStackCmd.Flags().String("description", "", "The description to display.")
		appstream_createStackCmd.Flags().String("display-name", "", "The stack name to display.")
		appstream_createStackCmd.Flags().String("embed-host-domains", "", "The domains where AppStream 2.0 streaming sessions can be embedded in an iframe.")
		appstream_createStackCmd.Flags().String("feedback-url", "", "The URL that users are redirected to after they click the Send Feedback link.")
		appstream_createStackCmd.Flags().String("name", "", "The name of the stack.")
		appstream_createStackCmd.Flags().String("redirect-url", "", "The URL that users are redirected to after their streaming session ends.")
		appstream_createStackCmd.Flags().String("storage-connectors", "", "The storage connectors to enable.")
		appstream_createStackCmd.Flags().String("streaming-experience-settings", "", "The streaming protocol you want your stack to prefer.")
		appstream_createStackCmd.Flags().String("tags", "", "The tags to associate with the stack.")
		appstream_createStackCmd.Flags().String("user-settings", "", "The actions that are enabled or disabled for users during their streaming sessions.")
		appstream_createStackCmd.MarkFlagRequired("name")
	})
	appstreamCmd.AddCommand(appstream_createStackCmd)
}
