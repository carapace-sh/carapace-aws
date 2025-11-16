package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getGlossaryTermCmd = &cobra.Command{
	Use:   "get-glossary-term",
	Short: "Gets a business glossary term in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getGlossaryTermCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_getGlossaryTermCmd).Standalone()

		datazone_getGlossaryTermCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which this business glossary term exists.")
		datazone_getGlossaryTermCmd.Flags().String("identifier", "", "The ID of the business glossary term.")
		datazone_getGlossaryTermCmd.MarkFlagRequired("domain-identifier")
		datazone_getGlossaryTermCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_getGlossaryTermCmd)
}
