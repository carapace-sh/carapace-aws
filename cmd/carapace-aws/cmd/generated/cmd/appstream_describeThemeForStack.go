package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_describeThemeForStackCmd = &cobra.Command{
	Use:   "describe-theme-for-stack",
	Short: "Retrieves a list that describes the theme for a specified stack.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_describeThemeForStackCmd).Standalone()

	appstream_describeThemeForStackCmd.Flags().String("stack-name", "", "The name of the stack for the theme.")
	appstream_describeThemeForStackCmd.MarkFlagRequired("stack-name")
	appstreamCmd.AddCommand(appstream_describeThemeForStackCmd)
}
