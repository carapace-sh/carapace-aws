package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_listTestsCmd = &cobra.Command{
	Use:   "list-tests",
	Short: "Gets information about tests in a given test suite.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_listTestsCmd).Standalone()

	devicefarm_listTestsCmd.Flags().String("arn", "", "The test suite's Amazon Resource Name (ARN).")
	devicefarm_listTestsCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which can be used to return the next set of items in the list.")
	devicefarm_listTestsCmd.MarkFlagRequired("arn")
	devicefarmCmd.AddCommand(devicefarm_listTestsCmd)
}
