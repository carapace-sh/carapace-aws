package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltower_getEnabledBaselineCmd = &cobra.Command{
	Use:   "get-enabled-baseline",
	Short: "Retrieve details of an `EnabledBaseline` resource by specifying its identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltower_getEnabledBaselineCmd).Standalone()

	controltower_getEnabledBaselineCmd.Flags().String("enabled-baseline-identifier", "", "Identifier of the `EnabledBaseline` resource to be retrieved, in ARN format.")
	controltower_getEnabledBaselineCmd.MarkFlagRequired("enabled-baseline-identifier")
	controltowerCmd.AddCommand(controltower_getEnabledBaselineCmd)
}
