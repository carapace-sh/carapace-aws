package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeMitigationActionCmd = &cobra.Command{
	Use:   "describe-mitigation-action",
	Short: "Gets information about a mitigation action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeMitigationActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_describeMitigationActionCmd).Standalone()

		iot_describeMitigationActionCmd.Flags().String("action-name", "", "The friendly name that uniquely identifies the mitigation action.")
		iot_describeMitigationActionCmd.MarkFlagRequired("action-name")
	})
	iotCmd.AddCommand(iot_describeMitigationActionCmd)
}
