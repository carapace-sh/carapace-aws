package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctions_describeMapRunCmd = &cobra.Command{
	Use:   "describe-map-run",
	Short: "Provides information about a Map Run's configuration, progress, and results.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctions_describeMapRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(stepfunctions_describeMapRunCmd).Standalone()

		stepfunctions_describeMapRunCmd.Flags().String("map-run-arn", "", "The Amazon Resource Name (ARN) that identifies a Map Run.")
		stepfunctions_describeMapRunCmd.MarkFlagRequired("map-run-arn")
	})
	stepfunctionsCmd.AddCommand(stepfunctions_describeMapRunCmd)
}
