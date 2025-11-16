package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehendmedical_detectPhiCmd = &cobra.Command{
	Use:   "detect-phi",
	Short: "Inspects the clinical text for protected health information (PHI) entities and returns the entity category, location, and confidence score for each entity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehendmedical_detectPhiCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehendmedical_detectPhiCmd).Standalone()

		comprehendmedical_detectPhiCmd.Flags().String("text", "", "A UTF-8 text string containing the clinical content being examined for PHI entities.")
		comprehendmedical_detectPhiCmd.MarkFlagRequired("text")
	})
	comprehendmedicalCmd.AddCommand(comprehendmedical_detectPhiCmd)
}
