package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_getSuiteCmd = &cobra.Command{
	Use:   "get-suite",
	Short: "Gets information about a suite.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_getSuiteCmd).Standalone()

	devicefarm_getSuiteCmd.Flags().String("arn", "", "The suite's ARN.")
	devicefarm_getSuiteCmd.MarkFlagRequired("arn")
	devicefarmCmd.AddCommand(devicefarm_getSuiteCmd)
}
