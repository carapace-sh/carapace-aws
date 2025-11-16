package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getUserProfileCmd = &cobra.Command{
	Use:   "get-user-profile",
	Short: "Gets a user profile in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getUserProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_getUserProfileCmd).Standalone()

		datazone_getUserProfileCmd.Flags().String("domain-identifier", "", "the ID of the Amazon DataZone domain the data portal of which you want to get.")
		datazone_getUserProfileCmd.Flags().String("type", "", "The type of the user profile.")
		datazone_getUserProfileCmd.Flags().String("user-identifier", "", "The identifier of the user for which you want to get the user profile.")
		datazone_getUserProfileCmd.MarkFlagRequired("domain-identifier")
		datazone_getUserProfileCmd.MarkFlagRequired("user-identifier")
	})
	datazoneCmd.AddCommand(datazone_getUserProfileCmd)
}
