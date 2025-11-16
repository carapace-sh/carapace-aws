package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mwaaCmd = &cobra.Command{
	Use:   "mwaa",
	Short: "Amazon Managed Workflows for Apache Airflow",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mwaaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mwaaCmd).Standalone()

	})
	rootCmd.AddCommand(mwaaCmd)
}
