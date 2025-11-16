package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_deleteProjectCmd = &cobra.Command{
	Use:   "delete-project",
	Short: "Deletes an existing DataBrew project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_deleteProjectCmd).Standalone()

	databrew_deleteProjectCmd.Flags().String("name", "", "The name of the project to be deleted.")
	databrew_deleteProjectCmd.MarkFlagRequired("name")
	databrewCmd.AddCommand(databrew_deleteProjectCmd)
}
