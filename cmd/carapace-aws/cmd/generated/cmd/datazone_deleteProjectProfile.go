package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_deleteProjectProfileCmd = &cobra.Command{
	Use:   "delete-project-profile",
	Short: "Deletes a project profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_deleteProjectProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_deleteProjectProfileCmd).Standalone()

		datazone_deleteProjectProfileCmd.Flags().String("domain-identifier", "", "The ID of the domain where a project profile is deleted.")
		datazone_deleteProjectProfileCmd.Flags().String("identifier", "", "The ID of the project profile that is deleted.")
		datazone_deleteProjectProfileCmd.MarkFlagRequired("domain-identifier")
		datazone_deleteProjectProfileCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_deleteProjectProfileCmd)
}
