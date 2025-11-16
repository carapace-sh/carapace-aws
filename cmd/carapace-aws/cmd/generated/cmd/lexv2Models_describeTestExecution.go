package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_describeTestExecutionCmd = &cobra.Command{
	Use:   "describe-test-execution",
	Short: "Gets metadata information about the test execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_describeTestExecutionCmd).Standalone()

	lexv2Models_describeTestExecutionCmd.Flags().String("test-execution-id", "", "The execution Id of the test set execution.")
	lexv2Models_describeTestExecutionCmd.MarkFlagRequired("test-execution-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_describeTestExecutionCmd)
}
