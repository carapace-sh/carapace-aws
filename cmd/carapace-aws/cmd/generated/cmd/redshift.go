package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftCmd = &cobra.Command{
	Use:   "redshift",
	Short: "Amazon Redshift",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftCmd).Standalone()

	})
	rootCmd.AddCommand(redshiftCmd)
}
