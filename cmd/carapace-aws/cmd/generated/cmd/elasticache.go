package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticacheCmd = &cobra.Command{
	Use:   "elasticache",
	Short: "Amazon ElastiCache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticacheCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticacheCmd).Standalone()

	})
	rootCmd.AddCommand(elasticacheCmd)
}
