package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcasesCmd = &cobra.Command{
	Use:   "connectcases",
	Short: "- [Cases actions](https://docs.aws.amazon.com/connect/latest/APIReference/API_Operations_Amazon_Connect_Cases.html)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcasesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcasesCmd).Standalone()

	})
	rootCmd.AddCommand(connectcasesCmd)
}
