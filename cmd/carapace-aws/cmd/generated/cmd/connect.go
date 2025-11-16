package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "- [Amazon Connect actions](https://docs.aws.amazon.com/connect/latest/APIReference/API_Operations_Amazon_Connect_Service.html)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectCmd).Standalone()

	})
	rootCmd.AddCommand(connectCmd)
}
