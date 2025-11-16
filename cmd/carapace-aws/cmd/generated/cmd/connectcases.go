package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcasesCmd = &cobra.Command{
	Use:   "connectcases",
	Short: "- [Cases actions](https://docs.aws.amazon.com/connect/latest/APIReference/API_Operations_Amazon_Connect_Cases.html)\n- [Cases data types](https://docs.aws.amazon.com/connect/latest/APIReference/API_Types_Amazon_Connect_Cases.html)\n\nWith Amazon Connect Cases, your agents can track and manage customer issues that require multiple interactions, follow-up tasks, and teams in your contact center.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcasesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcasesCmd).Standalone()

	})
	rootCmd.AddCommand(connectcasesCmd)
}
