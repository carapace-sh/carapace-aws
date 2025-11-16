package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidentlyCmd = &cobra.Command{
	Use:   "evidently",
	Short: "You can use Amazon CloudWatch Evidently to safely validate new features by serving them to a specified percentage of your users while you roll out the feature.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidentlyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(evidentlyCmd).Standalone()

	})
	rootCmd.AddCommand(evidentlyCmd)
}
