package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_createFormTypeCmd = &cobra.Command{
	Use:   "create-form-type",
	Short: "Creates a metadata form type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_createFormTypeCmd).Standalone()

	datazone_createFormTypeCmd.Flags().String("description", "", "The description of this Amazon DataZone metadata form type.")
	datazone_createFormTypeCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which this metadata form type is created.")
	datazone_createFormTypeCmd.Flags().String("model", "", "The model of this Amazon DataZone metadata form type.")
	datazone_createFormTypeCmd.Flags().String("name", "", "The name of this Amazon DataZone metadata form type.")
	datazone_createFormTypeCmd.Flags().String("owning-project-identifier", "", "The ID of the Amazon DataZone project that owns this metadata form type.")
	datazone_createFormTypeCmd.Flags().String("status", "", "The status of this Amazon DataZone metadata form type.")
	datazone_createFormTypeCmd.MarkFlagRequired("domain-identifier")
	datazone_createFormTypeCmd.MarkFlagRequired("model")
	datazone_createFormTypeCmd.MarkFlagRequired("name")
	datazone_createFormTypeCmd.MarkFlagRequired("owning-project-identifier")
	datazoneCmd.AddCommand(datazone_createFormTypeCmd)
}
