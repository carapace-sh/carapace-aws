package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrCmd = &cobra.Command{
	Use:   "emr",
	Short: "Amazon EMR is a web service that makes it easier to process large amounts of data efficiently.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrCmd).Standalone()

	rootCmd.AddCommand(emrCmd)
}
