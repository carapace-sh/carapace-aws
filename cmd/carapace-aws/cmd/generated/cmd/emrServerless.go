package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrServerlessCmd = &cobra.Command{
	Use:   "emr-serverless",
	Short: "Amazon EMR Serverless is a new deployment option for Amazon EMR.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrServerlessCmd).Standalone()

	rootCmd.AddCommand(emrServerlessCmd)
}
