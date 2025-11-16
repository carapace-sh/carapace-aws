package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_describeCaenrollmentPolicyCmd = &cobra.Command{
	Use:   "describe-caenrollment-policy",
	Short: "Retrieves detailed information about the certificate authority (CA) enrollment policy for the specified directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_describeCaenrollmentPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_describeCaenrollmentPolicyCmd).Standalone()

		ds_describeCaenrollmentPolicyCmd.Flags().String("directory-id", "", "The identifier of the directory for which to retrieve the CA enrollment policy information.")
		ds_describeCaenrollmentPolicyCmd.MarkFlagRequired("directory-id")
	})
	dsCmd.AddCommand(ds_describeCaenrollmentPolicyCmd)
}
