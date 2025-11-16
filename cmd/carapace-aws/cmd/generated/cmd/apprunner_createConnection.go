package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_createConnectionCmd = &cobra.Command{
	Use:   "create-connection",
	Short: "Create an App Runner connection resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_createConnectionCmd).Standalone()

	apprunner_createConnectionCmd.Flags().String("connection-name", "", "A name for the new connection.")
	apprunner_createConnectionCmd.Flags().String("provider-type", "", "The source repository provider.")
	apprunner_createConnectionCmd.Flags().String("tags", "", "A list of metadata items that you can associate with your connection resource.")
	apprunner_createConnectionCmd.MarkFlagRequired("connection-name")
	apprunner_createConnectionCmd.MarkFlagRequired("provider-type")
	apprunnerCmd.AddCommand(apprunner_createConnectionCmd)
}
