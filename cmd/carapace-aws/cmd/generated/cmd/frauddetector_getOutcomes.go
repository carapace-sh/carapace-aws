package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_getOutcomesCmd = &cobra.Command{
	Use:   "get-outcomes",
	Short: "Gets one or more outcomes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_getOutcomesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_getOutcomesCmd).Standalone()

		frauddetector_getOutcomesCmd.Flags().String("max-results", "", "The maximum number of objects to return for the request.")
		frauddetector_getOutcomesCmd.Flags().String("name", "", "The name of the outcome or outcomes to get.")
		frauddetector_getOutcomesCmd.Flags().String("next-token", "", "The next page token for the request.")
	})
	frauddetectorCmd.AddCommand(frauddetector_getOutcomesCmd)
}
