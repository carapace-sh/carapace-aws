package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_disableCaenrollmentPolicyCmd = &cobra.Command{
	Use:   "disable-caenrollment-policy",
	Short: "Disables the certificate authority (CA) enrollment policy for the specified directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_disableCaenrollmentPolicyCmd).Standalone()

	ds_disableCaenrollmentPolicyCmd.Flags().String("directory-id", "", "The identifier of the directory for which to disable the CA enrollment policy.")
	ds_disableCaenrollmentPolicyCmd.MarkFlagRequired("directory-id")
	dsCmd.AddCommand(ds_disableCaenrollmentPolicyCmd)
}
