package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltower_enableBaselineCmd = &cobra.Command{
	Use:   "enable-baseline",
	Short: "Enable (apply) a `Baseline` to a Target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltower_enableBaselineCmd).Standalone()

	controltower_enableBaselineCmd.Flags().String("baseline-identifier", "", "The ARN of the baseline to be enabled.")
	controltower_enableBaselineCmd.Flags().String("baseline-version", "", "The specific version to be enabled of the specified baseline.")
	controltower_enableBaselineCmd.Flags().String("parameters", "", "A list of `key-value` objects that specify enablement parameters, where `key` is a string and `value` is a document of any type.")
	controltower_enableBaselineCmd.Flags().String("tags", "", "Tags associated with input to `EnableBaseline`.")
	controltower_enableBaselineCmd.Flags().String("target-identifier", "", "The ARN of the target on which the baseline will be enabled.")
	controltower_enableBaselineCmd.MarkFlagRequired("baseline-identifier")
	controltower_enableBaselineCmd.MarkFlagRequired("baseline-version")
	controltower_enableBaselineCmd.MarkFlagRequired("target-identifier")
	controltowerCmd.AddCommand(controltower_enableBaselineCmd)
}
