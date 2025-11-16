package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getEnvironmentCredentialsCmd = &cobra.Command{
	Use:   "get-environment-credentials",
	Short: "Gets the credentials of an environment in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getEnvironmentCredentialsCmd).Standalone()

	datazone_getEnvironmentCredentialsCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which this environment and its credentials exist.")
	datazone_getEnvironmentCredentialsCmd.Flags().String("environment-identifier", "", "The ID of the environment whose credentials this operation gets.")
	datazone_getEnvironmentCredentialsCmd.MarkFlagRequired("domain-identifier")
	datazone_getEnvironmentCredentialsCmd.MarkFlagRequired("environment-identifier")
	datazoneCmd.AddCommand(datazone_getEnvironmentCredentialsCmd)
}
