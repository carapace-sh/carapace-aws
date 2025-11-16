package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getEnvironmentProfileCmd = &cobra.Command{
	Use:   "get-environment-profile",
	Short: "Gets an evinronment profile in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getEnvironmentProfileCmd).Standalone()

	datazone_getEnvironmentProfileCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which this environment profile exists.")
	datazone_getEnvironmentProfileCmd.Flags().String("identifier", "", "The ID of the environment profile.")
	datazone_getEnvironmentProfileCmd.MarkFlagRequired("domain-identifier")
	datazone_getEnvironmentProfileCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_getEnvironmentProfileCmd)
}
