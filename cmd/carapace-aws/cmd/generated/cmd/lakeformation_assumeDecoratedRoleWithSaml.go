package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_assumeDecoratedRoleWithSamlCmd = &cobra.Command{
	Use:   "assume-decorated-role-with-saml",
	Short: "Allows a caller to assume an IAM role decorated as the SAML user specified in the SAML assertion included in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_assumeDecoratedRoleWithSamlCmd).Standalone()

	lakeformation_assumeDecoratedRoleWithSamlCmd.Flags().String("duration-seconds", "", "The time period, between 900 and 43,200 seconds, for the timeout of the temporary credentials.")
	lakeformation_assumeDecoratedRoleWithSamlCmd.Flags().String("principal-arn", "", "The Amazon Resource Name (ARN) of the SAML provider in IAM that describes the IdP.")
	lakeformation_assumeDecoratedRoleWithSamlCmd.Flags().String("role-arn", "", "The role that represents an IAM principal whose scope down policy allows it to call credential vending APIs such as `GetTemporaryTableCredentials`.")
	lakeformation_assumeDecoratedRoleWithSamlCmd.Flags().String("samlassertion", "", "A SAML assertion consisting of an assertion statement for the user who needs temporary credentials.")
	lakeformation_assumeDecoratedRoleWithSamlCmd.MarkFlagRequired("principal-arn")
	lakeformation_assumeDecoratedRoleWithSamlCmd.MarkFlagRequired("role-arn")
	lakeformation_assumeDecoratedRoleWithSamlCmd.MarkFlagRequired("samlassertion")
	lakeformationCmd.AddCommand(lakeformation_assumeDecoratedRoleWithSamlCmd)
}
