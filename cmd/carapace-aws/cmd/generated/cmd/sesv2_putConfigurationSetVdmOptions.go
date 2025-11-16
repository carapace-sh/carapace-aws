package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_putConfigurationSetVdmOptionsCmd = &cobra.Command{
	Use:   "put-configuration-set-vdm-options",
	Short: "Specify VDM preferences for email that you send using the configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_putConfigurationSetVdmOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_putConfigurationSetVdmOptionsCmd).Standalone()

		sesv2_putConfigurationSetVdmOptionsCmd.Flags().String("configuration-set-name", "", "The name of the configuration set.")
		sesv2_putConfigurationSetVdmOptionsCmd.Flags().String("vdm-options", "", "The VDM options to apply to the configuration set.")
		sesv2_putConfigurationSetVdmOptionsCmd.MarkFlagRequired("configuration-set-name")
	})
	sesv2Cmd.AddCommand(sesv2_putConfigurationSetVdmOptionsCmd)
}
