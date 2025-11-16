package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_deleteHsmConfigurationCmd = &cobra.Command{
	Use:   "delete-hsm-configuration",
	Short: "Deletes the specified Amazon Redshift HSM configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_deleteHsmConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_deleteHsmConfigurationCmd).Standalone()

		redshift_deleteHsmConfigurationCmd.Flags().String("hsm-configuration-identifier", "", "The identifier of the Amazon Redshift HSM configuration to be deleted.")
		redshift_deleteHsmConfigurationCmd.MarkFlagRequired("hsm-configuration-identifier")
	})
	redshiftCmd.AddCommand(redshift_deleteHsmConfigurationCmd)
}
