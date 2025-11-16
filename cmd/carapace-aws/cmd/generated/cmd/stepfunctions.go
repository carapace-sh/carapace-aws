package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctionsCmd = &cobra.Command{
	Use:   "stepfunctions",
	Short: "Step Functions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(stepfunctionsCmd).Standalone()

	})
	rootCmd.AddCommand(stepfunctionsCmd)
}
