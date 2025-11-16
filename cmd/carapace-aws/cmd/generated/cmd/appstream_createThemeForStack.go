package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_createThemeForStackCmd = &cobra.Command{
	Use:   "create-theme-for-stack",
	Short: "Creates custom branding that customizes the appearance of the streaming application catalog page.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_createThemeForStackCmd).Standalone()

	appstream_createThemeForStackCmd.Flags().String("favicon-s3-location", "", "The S3 location of the favicon.")
	appstream_createThemeForStackCmd.Flags().String("footer-links", "", "The links that are displayed in the footer of the streaming application catalog page.")
	appstream_createThemeForStackCmd.Flags().String("organization-logo-s3-location", "", "The organization logo that appears on the streaming application catalog page.")
	appstream_createThemeForStackCmd.Flags().String("stack-name", "", "The name of the stack for the theme.")
	appstream_createThemeForStackCmd.Flags().String("theme-styling", "", "The color theme that is applied to website links, text, and buttons.")
	appstream_createThemeForStackCmd.Flags().String("title-text", "", "The title that is displayed at the top of the browser tab during users' application streaming sessions.")
	appstream_createThemeForStackCmd.MarkFlagRequired("favicon-s3-location")
	appstream_createThemeForStackCmd.MarkFlagRequired("organization-logo-s3-location")
	appstream_createThemeForStackCmd.MarkFlagRequired("stack-name")
	appstream_createThemeForStackCmd.MarkFlagRequired("theme-styling")
	appstream_createThemeForStackCmd.MarkFlagRequired("title-text")
	appstreamCmd.AddCommand(appstream_createThemeForStackCmd)
}
