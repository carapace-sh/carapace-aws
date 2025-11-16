package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourcegroupstaggingapiCmd = &cobra.Command{
	Use:   "resourcegroupstaggingapi",
	Short: "Resource Groups Tagging API",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourcegroupstaggingapiCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourcegroupstaggingapiCmd).Standalone()

	})
	rootCmd.AddCommand(resourcegroupstaggingapiCmd)
}
