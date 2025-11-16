package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_putResourcePolicyCmd = &cobra.Command{
	Use:   "put-resource-policy",
	Short: "Sets the Data Catalog resource policy for access control.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_putResourcePolicyCmd).Standalone()

	glue_putResourcePolicyCmd.Flags().String("enable-hybrid", "", "If `'TRUE'`, indicates that you are using both methods to grant cross-account access to Data Catalog resources:")
	glue_putResourcePolicyCmd.Flags().String("policy-exists-condition", "", "A value of `MUST_EXIST` is used to update a policy.")
	glue_putResourcePolicyCmd.Flags().String("policy-hash-condition", "", "The hash value returned when the previous policy was set using `PutResourcePolicy`.")
	glue_putResourcePolicyCmd.Flags().String("policy-in-json", "", "Contains the policy document to set, in JSON format.")
	glue_putResourcePolicyCmd.Flags().String("resource-arn", "", "Do not use.")
	glue_putResourcePolicyCmd.MarkFlagRequired("policy-in-json")
	glueCmd.AddCommand(glue_putResourcePolicyCmd)
}
