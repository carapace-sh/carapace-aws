package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_deleteConfigurationSetTrackingOptionsCmd = &cobra.Command{
	Use:   "delete-configuration-set-tracking-options",
	Short: "Deletes an association between a configuration set and a custom domain for open and click event tracking.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_deleteConfigurationSetTrackingOptionsCmd).Standalone()

	ses_deleteConfigurationSetTrackingOptionsCmd.Flags().String("configuration-set-name", "", "The name of the configuration set.")
	ses_deleteConfigurationSetTrackingOptionsCmd.MarkFlagRequired("configuration-set-name")
	sesCmd.AddCommand(ses_deleteConfigurationSetTrackingOptionsCmd)
}
