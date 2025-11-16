package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehendmedical_detectEntitiesCmd = &cobra.Command{
	Use:   "detect-entities",
	Short: "The `DetectEntities` operation is deprecated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehendmedical_detectEntitiesCmd).Standalone()

	comprehendmedical_detectEntitiesCmd.Flags().String("text", "", "A UTF-8 text string containing the clinical content being examined for entities.")
	comprehendmedical_detectEntitiesCmd.MarkFlagRequired("text")
	comprehendmedicalCmd.AddCommand(comprehendmedical_detectEntitiesCmd)
}
