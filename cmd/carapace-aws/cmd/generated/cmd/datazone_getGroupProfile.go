package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getGroupProfileCmd = &cobra.Command{
	Use:   "get-group-profile",
	Short: "Gets a group profile in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getGroupProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_getGroupProfileCmd).Standalone()

		datazone_getGroupProfileCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain in which the group profile exists.")
		datazone_getGroupProfileCmd.Flags().String("group-identifier", "", "The identifier of the group profile.")
		datazone_getGroupProfileCmd.MarkFlagRequired("domain-identifier")
		datazone_getGroupProfileCmd.MarkFlagRequired("group-identifier")
	})
	datazoneCmd.AddCommand(datazone_getGroupProfileCmd)
}
