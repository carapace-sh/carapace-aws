package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_createConfigurationSetTrackingOptionsCmd = &cobra.Command{
	Use:   "create-configuration-set-tracking-options",
	Short: "Creates an association between a configuration set and a custom domain for open and click event tracking.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_createConfigurationSetTrackingOptionsCmd).Standalone()

	ses_createConfigurationSetTrackingOptionsCmd.Flags().String("configuration-set-name", "", "The name of the configuration set that the tracking options should be associated with.")
	ses_createConfigurationSetTrackingOptionsCmd.Flags().String("tracking-options", "", "")
	ses_createConfigurationSetTrackingOptionsCmd.MarkFlagRequired("configuration-set-name")
	ses_createConfigurationSetTrackingOptionsCmd.MarkFlagRequired("tracking-options")
	sesCmd.AddCommand(ses_createConfigurationSetTrackingOptionsCmd)
}
