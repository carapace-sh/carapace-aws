package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_getLabelsCmd = &cobra.Command{
	Use:   "get-labels",
	Short: "Gets all labels or a specific label if name is provided.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_getLabelsCmd).Standalone()

	frauddetector_getLabelsCmd.Flags().String("max-results", "", "The maximum number of objects to return for the request.")
	frauddetector_getLabelsCmd.Flags().String("name", "", "The name of the label or labels to get.")
	frauddetector_getLabelsCmd.Flags().String("next-token", "", "The next token for the subsequent request.")
	frauddetectorCmd.AddCommand(frauddetector_getLabelsCmd)
}
