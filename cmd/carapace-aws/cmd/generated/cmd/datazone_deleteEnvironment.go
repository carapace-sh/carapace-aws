package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_deleteEnvironmentCmd = &cobra.Command{
	Use:   "delete-environment",
	Short: "Deletes an environment in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_deleteEnvironmentCmd).Standalone()

	datazone_deleteEnvironmentCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which the environment is deleted.")
	datazone_deleteEnvironmentCmd.Flags().String("identifier", "", "The identifier of the environment that is to be deleted.")
	datazone_deleteEnvironmentCmd.MarkFlagRequired("domain-identifier")
	datazone_deleteEnvironmentCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_deleteEnvironmentCmd)
}
