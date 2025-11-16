package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_getBrowserSettingsCmd = &cobra.Command{
	Use:   "get-browser-settings",
	Short: "Gets browser settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_getBrowserSettingsCmd).Standalone()

	workspacesWeb_getBrowserSettingsCmd.Flags().String("browser-settings-arn", "", "The ARN of the browser settings.")
	workspacesWeb_getBrowserSettingsCmd.MarkFlagRequired("browser-settings-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_getBrowserSettingsCmd)
}
