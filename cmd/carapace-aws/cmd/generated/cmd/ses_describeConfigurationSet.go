package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_describeConfigurationSetCmd = &cobra.Command{
	Use:   "describe-configuration-set",
	Short: "Returns the details of the specified configuration set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_describeConfigurationSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_describeConfigurationSetCmd).Standalone()

		ses_describeConfigurationSetCmd.Flags().String("configuration-set-attribute-names", "", "A list of configuration set attributes to return.")
		ses_describeConfigurationSetCmd.Flags().String("configuration-set-name", "", "The name of the configuration set to describe.")
		ses_describeConfigurationSetCmd.MarkFlagRequired("configuration-set-name")
	})
	sesCmd.AddCommand(ses_describeConfigurationSetCmd)
}
