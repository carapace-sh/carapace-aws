package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_testConnectionCmd = &cobra.Command{
	Use:   "test-connection",
	Short: "Tests a connection to a service to validate the service credentials that you provide.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_testConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_testConnectionCmd).Standalone()

		glue_testConnectionCmd.Flags().String("catalog-id", "", "The catalog ID where the connection resides.")
		glue_testConnectionCmd.Flags().String("connection-name", "", "Optional.")
		glue_testConnectionCmd.Flags().String("test-connection-input", "", "A structure that is used to specify testing a connection to a service.")
	})
	glueCmd.AddCommand(glue_testConnectionCmd)
}
