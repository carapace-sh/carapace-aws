package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltower_disableBaselineCmd = &cobra.Command{
	Use:   "disable-baseline",
	Short: "Disable an `EnabledBaseline` resource on the specified Target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltower_disableBaselineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(controltower_disableBaselineCmd).Standalone()

		controltower_disableBaselineCmd.Flags().String("enabled-baseline-identifier", "", "Identifier of the `EnabledBaseline` resource to be deactivated, in ARN format.")
		controltower_disableBaselineCmd.MarkFlagRequired("enabled-baseline-identifier")
	})
	controltowerCmd.AddCommand(controltower_disableBaselineCmd)
}
