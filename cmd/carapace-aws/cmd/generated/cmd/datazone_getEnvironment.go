package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getEnvironmentCmd = &cobra.Command{
	Use:   "get-environment",
	Short: "Gets an Amazon DataZone environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getEnvironmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_getEnvironmentCmd).Standalone()

		datazone_getEnvironmentCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain where the environment exists.")
		datazone_getEnvironmentCmd.Flags().String("identifier", "", "The ID of the Amazon DataZone environment.")
		datazone_getEnvironmentCmd.MarkFlagRequired("domain-identifier")
		datazone_getEnvironmentCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_getEnvironmentCmd)
}
