package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_describeTypeRegistrationCmd = &cobra.Command{
	Use:   "describe-type-registration",
	Short: "Returns information about an extension's registration, including its current status and type and version identifiers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_describeTypeRegistrationCmd).Standalone()

	cloudformation_describeTypeRegistrationCmd.Flags().String("registration-token", "", "The identifier for this registration request.")
	cloudformation_describeTypeRegistrationCmd.MarkFlagRequired("registration-token")
	cloudformationCmd.AddCommand(cloudformation_describeTypeRegistrationCmd)
}
