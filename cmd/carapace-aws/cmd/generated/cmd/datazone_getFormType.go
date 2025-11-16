package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getFormTypeCmd = &cobra.Command{
	Use:   "get-form-type",
	Short: "Gets a metadata form type in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getFormTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_getFormTypeCmd).Standalone()

		datazone_getFormTypeCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which this metadata form type exists.")
		datazone_getFormTypeCmd.Flags().String("form-type-identifier", "", "The ID of the metadata form type.")
		datazone_getFormTypeCmd.Flags().String("revision", "", "The revision of this metadata form type.")
		datazone_getFormTypeCmd.MarkFlagRequired("domain-identifier")
		datazone_getFormTypeCmd.MarkFlagRequired("form-type-identifier")
	})
	datazoneCmd.AddCommand(datazone_getFormTypeCmd)
}
