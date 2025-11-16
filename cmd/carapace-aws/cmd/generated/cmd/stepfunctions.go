package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stepfunctionsCmd = &cobra.Command{
	Use:   "stepfunctions",
	Short: "Step Functions\n\nWith Step Functions, you can create workflows, also called *state machines*, to build distributed applications, automate processes, orchestrate microservices, and create data and machine learning pipelines.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stepfunctionsCmd).Standalone()

	rootCmd.AddCommand(stepfunctionsCmd)
}
