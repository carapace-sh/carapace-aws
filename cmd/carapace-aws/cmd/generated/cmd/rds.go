package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rdsCmd = &cobra.Command{
	Use:   "rds",
	Short: "Amazon Relational Database Service\n\nAmazon Relational Database Service (Amazon RDS) is a web service that makes it easier to set up, operate, and scale a relational database in the cloud.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rdsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rdsCmd).Standalone()

	})
	rootCmd.AddCommand(rdsCmd)
}
