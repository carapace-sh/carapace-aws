package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getGlossaryCmd = &cobra.Command{
	Use:   "get-glossary",
	Short: "Gets a business glossary in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getGlossaryCmd).Standalone()

	datazone_getGlossaryCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which this business glossary exists.")
	datazone_getGlossaryCmd.Flags().String("identifier", "", "The ID of the business glossary.")
	datazone_getGlossaryCmd.MarkFlagRequired("domain-identifier")
	datazone_getGlossaryCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_getGlossaryCmd)
}
