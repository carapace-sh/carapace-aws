package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_deleteFormTypeCmd = &cobra.Command{
	Use:   "delete-form-type",
	Short: "Deletes and metadata form type in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_deleteFormTypeCmd).Standalone()

	datazone_deleteFormTypeCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which the metadata form type is deleted.")
	datazone_deleteFormTypeCmd.Flags().String("form-type-identifier", "", "The ID of the metadata form type that is deleted.")
	datazone_deleteFormTypeCmd.MarkFlagRequired("domain-identifier")
	datazone_deleteFormTypeCmd.MarkFlagRequired("form-type-identifier")
	datazoneCmd.AddCommand(datazone_deleteFormTypeCmd)
}
