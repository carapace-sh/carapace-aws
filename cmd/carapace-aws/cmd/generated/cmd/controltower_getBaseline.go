package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltower_getBaselineCmd = &cobra.Command{
	Use:   "get-baseline",
	Short: "Retrieve details about an existing `Baseline` resource by specifying its identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltower_getBaselineCmd).Standalone()

	controltower_getBaselineCmd.Flags().String("baseline-identifier", "", "The ARN of the `Baseline` resource to be retrieved.")
	controltower_getBaselineCmd.MarkFlagRequired("baseline-identifier")
	controltowerCmd.AddCommand(controltower_getBaselineCmd)
}
