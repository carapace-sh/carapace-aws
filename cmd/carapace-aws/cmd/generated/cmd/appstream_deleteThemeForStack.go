package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_deleteThemeForStackCmd = &cobra.Command{
	Use:   "delete-theme-for-stack",
	Short: "Deletes custom branding that customizes the appearance of the streaming application catalog page.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_deleteThemeForStackCmd).Standalone()

	appstream_deleteThemeForStackCmd.Flags().String("stack-name", "", "The name of the stack for the theme.")
	appstream_deleteThemeForStackCmd.MarkFlagRequired("stack-name")
	appstreamCmd.AddCommand(appstream_deleteThemeForStackCmd)
}
