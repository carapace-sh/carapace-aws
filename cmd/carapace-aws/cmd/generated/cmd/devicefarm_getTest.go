package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_getTestCmd = &cobra.Command{
	Use:   "get-test",
	Short: "Gets information about a test.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_getTestCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_getTestCmd).Standalone()

		devicefarm_getTestCmd.Flags().String("arn", "", "The test's ARN.")
		devicefarm_getTestCmd.MarkFlagRequired("arn")
	})
	devicefarmCmd.AddCommand(devicefarm_getTestCmd)
}
