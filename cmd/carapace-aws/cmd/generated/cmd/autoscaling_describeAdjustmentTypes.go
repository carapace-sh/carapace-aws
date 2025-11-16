package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_describeAdjustmentTypesCmd = &cobra.Command{
	Use:   "describe-adjustment-types",
	Short: "Describes the available adjustment types for step scaling and simple scaling policies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_describeAdjustmentTypesCmd).Standalone()

	autoscalingCmd.AddCommand(autoscaling_describeAdjustmentTypesCmd)
}
