package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_createGroupProfileCmd = &cobra.Command{
	Use:   "create-group-profile",
	Short: "Creates a group profile in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_createGroupProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_createGroupProfileCmd).Standalone()

		datazone_createGroupProfileCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
		datazone_createGroupProfileCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain in which the group profile is created.")
		datazone_createGroupProfileCmd.Flags().String("group-identifier", "", "The identifier of the group for which the group profile is created.")
		datazone_createGroupProfileCmd.MarkFlagRequired("domain-identifier")
		datazone_createGroupProfileCmd.MarkFlagRequired("group-identifier")
	})
	datazoneCmd.AddCommand(datazone_createGroupProfileCmd)
}
