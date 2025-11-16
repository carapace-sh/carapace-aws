package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rdsDataCmd = &cobra.Command{
	Use:   "rds-data",
	Short: "RDS Data API\n\nAmazon RDS provides an HTTP endpoint to run SQL statements on an Amazon Aurora DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rdsDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rdsDataCmd).Standalone()

	})
	rootCmd.AddCommand(rdsDataCmd)
}
