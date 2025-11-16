package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerlessCmd = &cobra.Command{
	Use:   "redshift-serverless",
	Short: "This is an interface reference for Amazon Redshift Serverless.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerlessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftServerlessCmd).Standalone()

	})
	rootCmd.AddCommand(redshiftServerlessCmd)
}
