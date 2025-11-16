package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftDataCmd = &cobra.Command{
	Use:   "redshift-data",
	Short: "You can use the Amazon Redshift Data API to run queries on Amazon Redshift tables.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftDataCmd).Standalone()

	})
	rootCmd.AddCommand(redshiftDataCmd)
}
