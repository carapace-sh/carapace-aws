package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getProjectProfileCmd = &cobra.Command{
	Use:   "get-project-profile",
	Short: "The details of the project profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getProjectProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_getProjectProfileCmd).Standalone()

		datazone_getProjectProfileCmd.Flags().String("domain-identifier", "", "The ID of the domain.")
		datazone_getProjectProfileCmd.Flags().String("identifier", "", "The ID of the project profile.")
		datazone_getProjectProfileCmd.MarkFlagRequired("domain-identifier")
		datazone_getProjectProfileCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_getProjectProfileCmd)
}
