package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_updateStackCmd = &cobra.Command{
	Use:   "update-stack",
	Short: "Updates the specified fields for the specified stack.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_updateStackCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_updateStackCmd).Standalone()

		appstream_updateStackCmd.Flags().String("access-endpoints", "", "The list of interface VPC endpoint (interface endpoint) objects.")
		appstream_updateStackCmd.Flags().String("application-settings", "", "The persistent application settings for users of a stack.")
		appstream_updateStackCmd.Flags().String("attributes-to-delete", "", "The stack attributes to delete.")
		appstream_updateStackCmd.Flags().Bool("delete-storage-connectors", false, "Deletes the storage connectors currently enabled for the stack.")
		appstream_updateStackCmd.Flags().String("description", "", "The description to display.")
		appstream_updateStackCmd.Flags().String("display-name", "", "The stack name to display.")
		appstream_updateStackCmd.Flags().String("embed-host-domains", "", "The domains where AppStream 2.0 streaming sessions can be embedded in an iframe.")
		appstream_updateStackCmd.Flags().String("feedback-url", "", "The URL that users are redirected to after they choose the Send Feedback link.")
		appstream_updateStackCmd.Flags().String("name", "", "The name of the stack.")
		appstream_updateStackCmd.Flags().Bool("no-delete-storage-connectors", false, "Deletes the storage connectors currently enabled for the stack.")
		appstream_updateStackCmd.Flags().String("redirect-url", "", "The URL that users are redirected to after their streaming session ends.")
		appstream_updateStackCmd.Flags().String("storage-connectors", "", "The storage connectors to enable.")
		appstream_updateStackCmd.Flags().String("streaming-experience-settings", "", "The streaming protocol you want your stack to prefer.")
		appstream_updateStackCmd.Flags().String("user-settings", "", "The actions that are enabled or disabled for users during their streaming sessions.")
		appstream_updateStackCmd.MarkFlagRequired("name")
		appstream_updateStackCmd.Flag("no-delete-storage-connectors").Hidden = true
	})
	appstreamCmd.AddCommand(appstream_updateStackCmd)
}
