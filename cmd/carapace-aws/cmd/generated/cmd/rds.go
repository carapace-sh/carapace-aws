package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rdsCmd = &cobra.Command{
	Use:   "rds",
	Short: "Amazon Relational Database Service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rdsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rdsCmd).Standalone()

	})
	rootCmd.AddCommand(rdsCmd)
}
