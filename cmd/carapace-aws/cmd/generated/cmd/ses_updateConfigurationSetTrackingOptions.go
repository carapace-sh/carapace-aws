package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_updateConfigurationSetTrackingOptionsCmd = &cobra.Command{
	Use:   "update-configuration-set-tracking-options",
	Short: "Modifies an association between a configuration set and a custom domain for open and click event tracking.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_updateConfigurationSetTrackingOptionsCmd).Standalone()

	ses_updateConfigurationSetTrackingOptionsCmd.Flags().String("configuration-set-name", "", "The name of the configuration set.")
	ses_updateConfigurationSetTrackingOptionsCmd.Flags().String("tracking-options", "", "")
	ses_updateConfigurationSetTrackingOptionsCmd.MarkFlagRequired("configuration-set-name")
	ses_updateConfigurationSetTrackingOptionsCmd.MarkFlagRequired("tracking-options")
	sesCmd.AddCommand(ses_updateConfigurationSetTrackingOptionsCmd)
}
