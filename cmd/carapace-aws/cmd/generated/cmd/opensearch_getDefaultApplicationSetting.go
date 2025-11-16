package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_getDefaultApplicationSettingCmd = &cobra.Command{
	Use:   "get-default-application-setting",
	Short: "Gets the ARN of the current default application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_getDefaultApplicationSettingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_getDefaultApplicationSettingCmd).Standalone()

	})
	opensearchCmd.AddCommand(opensearch_getDefaultApplicationSettingCmd)
}
