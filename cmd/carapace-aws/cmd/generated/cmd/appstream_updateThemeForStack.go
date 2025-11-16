package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_updateThemeForStackCmd = &cobra.Command{
	Use:   "update-theme-for-stack",
	Short: "Updates custom branding that customizes the appearance of the streaming application catalog page.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_updateThemeForStackCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_updateThemeForStackCmd).Standalone()

		appstream_updateThemeForStackCmd.Flags().String("attributes-to-delete", "", "The attributes to delete.")
		appstream_updateThemeForStackCmd.Flags().String("favicon-s3-location", "", "The S3 location of the favicon.")
		appstream_updateThemeForStackCmd.Flags().String("footer-links", "", "The links that are displayed in the footer of the streaming application catalog page.")
		appstream_updateThemeForStackCmd.Flags().String("organization-logo-s3-location", "", "The organization logo that appears on the streaming application catalog page.")
		appstream_updateThemeForStackCmd.Flags().String("stack-name", "", "The name of the stack for the theme.")
		appstream_updateThemeForStackCmd.Flags().String("state", "", "Specifies whether custom branding should be applied to catalog page or not.")
		appstream_updateThemeForStackCmd.Flags().String("theme-styling", "", "The color theme that is applied to website links, text, and buttons.")
		appstream_updateThemeForStackCmd.Flags().String("title-text", "", "The title that is displayed at the top of the browser tab during users' application streaming sessions.")
		appstream_updateThemeForStackCmd.MarkFlagRequired("stack-name")
	})
	appstreamCmd.AddCommand(appstream_updateThemeForStackCmd)
}
