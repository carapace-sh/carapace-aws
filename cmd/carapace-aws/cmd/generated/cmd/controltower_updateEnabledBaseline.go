package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltower_updateEnabledBaselineCmd = &cobra.Command{
	Use:   "update-enabled-baseline",
	Short: "Updates an `EnabledBaseline` resource's applied parameters or version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltower_updateEnabledBaselineCmd).Standalone()

	controltower_updateEnabledBaselineCmd.Flags().String("baseline-version", "", "Specifies the new `Baseline` version, to which the `EnabledBaseline` should be updated.")
	controltower_updateEnabledBaselineCmd.Flags().String("enabled-baseline-identifier", "", "Specifies the `EnabledBaseline` resource to be updated.")
	controltower_updateEnabledBaselineCmd.Flags().String("parameters", "", "Parameters to apply when making an update.")
	controltower_updateEnabledBaselineCmd.MarkFlagRequired("baseline-version")
	controltower_updateEnabledBaselineCmd.MarkFlagRequired("enabled-baseline-identifier")
	controltowerCmd.AddCommand(controltower_updateEnabledBaselineCmd)
}
