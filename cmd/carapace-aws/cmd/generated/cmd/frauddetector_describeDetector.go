package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_describeDetectorCmd = &cobra.Command{
	Use:   "describe-detector",
	Short: "Gets all versions for a specified detector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_describeDetectorCmd).Standalone()

	frauddetector_describeDetectorCmd.Flags().String("detector-id", "", "The detector ID.")
	frauddetector_describeDetectorCmd.Flags().String("max-results", "", "The maximum number of results to return for the request.")
	frauddetector_describeDetectorCmd.Flags().String("next-token", "", "The next token from the previous response.")
	frauddetector_describeDetectorCmd.MarkFlagRequired("detector-id")
	frauddetectorCmd.AddCommand(frauddetector_describeDetectorCmd)
}
