package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getProjectCmd = &cobra.Command{
	Use:   "get-project",
	Short: "Gets a project in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getProjectCmd).Standalone()

	datazone_getProjectCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which the project exists.")
	datazone_getProjectCmd.Flags().String("identifier", "", "The ID of the project.")
	datazone_getProjectCmd.MarkFlagRequired("domain-identifier")
	datazone_getProjectCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_getProjectCmd)
}
