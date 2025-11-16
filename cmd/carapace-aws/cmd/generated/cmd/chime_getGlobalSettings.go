package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_getGlobalSettingsCmd = &cobra.Command{
	Use:   "get-global-settings",
	Short: "Retrieves global settings for the administrator's AWS account, such as Amazon Chime Business Calling and Amazon Chime Voice Connector settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_getGlobalSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_getGlobalSettingsCmd).Standalone()

	})
	chimeCmd.AddCommand(chime_getGlobalSettingsCmd)
}
