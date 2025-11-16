package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_createConnectionCmd = &cobra.Command{
	Use:   "create-connection",
	Short: "Creates a new connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_createConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_createConnectionCmd).Standalone()

		datazone_createConnectionCmd.Flags().String("aws-location", "", "The location where the connection is created.")
		datazone_createConnectionCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
		datazone_createConnectionCmd.Flags().String("description", "", "A connection description.")
		datazone_createConnectionCmd.Flags().String("domain-identifier", "", "The ID of the domain where the connection is created.")
		datazone_createConnectionCmd.Flags().Bool("enable-trusted-identity-propagation", false, "Specifies whether the trusted identity propagation is enabled.")
		datazone_createConnectionCmd.Flags().String("environment-identifier", "", "The ID of the environment where the connection is created.")
		datazone_createConnectionCmd.Flags().String("name", "", "The connection name.")
		datazone_createConnectionCmd.Flags().Bool("no-enable-trusted-identity-propagation", false, "Specifies whether the trusted identity propagation is enabled.")
		datazone_createConnectionCmd.Flags().String("props", "", "The connection props.")
		datazone_createConnectionCmd.Flags().String("scope", "", "The scope of the connection.")
		datazone_createConnectionCmd.MarkFlagRequired("domain-identifier")
		datazone_createConnectionCmd.MarkFlagRequired("name")
		datazone_createConnectionCmd.Flag("no-enable-trusted-identity-propagation").Hidden = true
	})
	datazoneCmd.AddCommand(datazone_createConnectionCmd)
}
