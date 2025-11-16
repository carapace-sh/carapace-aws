package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_enableCaenrollmentPolicyCmd = &cobra.Command{
	Use:   "enable-caenrollment-policy",
	Short: "Enables certificate authority (CA) enrollment policy for the specified directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_enableCaenrollmentPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_enableCaenrollmentPolicyCmd).Standalone()

		ds_enableCaenrollmentPolicyCmd.Flags().String("directory-id", "", "The identifier of the directory for which to enable the CA enrollment policy.")
		ds_enableCaenrollmentPolicyCmd.Flags().String("pca-connector-arn", "", "The Amazon Resource Name (ARN) of the Private Certificate Authority (PCA) connector to use for automatic certificate enrollment.")
		ds_enableCaenrollmentPolicyCmd.MarkFlagRequired("directory-id")
		ds_enableCaenrollmentPolicyCmd.MarkFlagRequired("pca-connector-arn")
	})
	dsCmd.AddCommand(ds_enableCaenrollmentPolicyCmd)
}
