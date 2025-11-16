package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_describeLifecycleHookTypesCmd = &cobra.Command{
	Use:   "describe-lifecycle-hook-types",
	Short: "Describes the available types of lifecycle hooks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_describeLifecycleHookTypesCmd).Standalone()

	autoscalingCmd.AddCommand(autoscaling_describeLifecycleHookTypesCmd)
}
