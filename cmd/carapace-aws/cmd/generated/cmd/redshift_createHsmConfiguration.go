package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_createHsmConfigurationCmd = &cobra.Command{
	Use:   "create-hsm-configuration",
	Short: "Creates an HSM configuration that contains the information required by an Amazon Redshift cluster to store and use database encryption keys in a Hardware Security Module (HSM).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_createHsmConfigurationCmd).Standalone()

	redshift_createHsmConfigurationCmd.Flags().String("description", "", "A text description of the HSM configuration to be created.")
	redshift_createHsmConfigurationCmd.Flags().String("hsm-configuration-identifier", "", "The identifier to be assigned to the new Amazon Redshift HSM configuration.")
	redshift_createHsmConfigurationCmd.Flags().String("hsm-ip-address", "", "The IP address that the Amazon Redshift cluster must use to access the HSM.")
	redshift_createHsmConfigurationCmd.Flags().String("hsm-partition-name", "", "The name of the partition in the HSM where the Amazon Redshift clusters will store their database encryption keys.")
	redshift_createHsmConfigurationCmd.Flags().String("hsm-partition-password", "", "The password required to access the HSM partition.")
	redshift_createHsmConfigurationCmd.Flags().String("hsm-server-public-certificate", "", "The HSMs public certificate file.")
	redshift_createHsmConfigurationCmd.Flags().String("tags", "", "A list of tag instances.")
	redshift_createHsmConfigurationCmd.MarkFlagRequired("description")
	redshift_createHsmConfigurationCmd.MarkFlagRequired("hsm-configuration-identifier")
	redshift_createHsmConfigurationCmd.MarkFlagRequired("hsm-ip-address")
	redshift_createHsmConfigurationCmd.MarkFlagRequired("hsm-partition-name")
	redshift_createHsmConfigurationCmd.MarkFlagRequired("hsm-partition-password")
	redshift_createHsmConfigurationCmd.MarkFlagRequired("hsm-server-public-certificate")
	redshiftCmd.AddCommand(redshift_createHsmConfigurationCmd)
}
