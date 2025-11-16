package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltower_resetEnabledBaselineCmd = &cobra.Command{
	Use:   "reset-enabled-baseline",
	Short: "Re-enables an `EnabledBaseline` resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltower_resetEnabledBaselineCmd).Standalone()

	controltower_resetEnabledBaselineCmd.Flags().String("enabled-baseline-identifier", "", "Specifies the ID of the `EnabledBaseline` resource to be re-enabled, in ARN format.")
	controltower_resetEnabledBaselineCmd.MarkFlagRequired("enabled-baseline-identifier")
	controltowerCmd.AddCommand(controltower_resetEnabledBaselineCmd)
}
