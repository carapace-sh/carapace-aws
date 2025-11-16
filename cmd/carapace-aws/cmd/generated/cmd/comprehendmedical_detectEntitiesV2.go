package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehendmedical_detectEntitiesV2Cmd = &cobra.Command{
	Use:   "detect-entities-v2",
	Short: "Inspects the clinical text for a variety of medical entities and returns specific information about them such as entity category, location, and confidence score on that information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehendmedical_detectEntitiesV2Cmd).Standalone()

	comprehendmedical_detectEntitiesV2Cmd.Flags().String("text", "", "A UTF-8 string containing the clinical content being examined for entities.")
	comprehendmedical_detectEntitiesV2Cmd.MarkFlagRequired("text")
	comprehendmedicalCmd.AddCommand(comprehendmedical_detectEntitiesV2Cmd)
}
