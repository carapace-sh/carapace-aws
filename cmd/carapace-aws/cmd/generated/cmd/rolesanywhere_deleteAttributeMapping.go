package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_deleteAttributeMappingCmd = &cobra.Command{
	Use:   "delete-attribute-mapping",
	Short: "Delete an entry from the attribute mapping rules enforced by a given profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_deleteAttributeMappingCmd).Standalone()

	rolesanywhere_deleteAttributeMappingCmd.Flags().String("certificate-field", "", "Fields (x509Subject, x509Issuer and x509SAN) within X.509 certificates.")
	rolesanywhere_deleteAttributeMappingCmd.Flags().String("profile-id", "", "The unique identifier of the profile.")
	rolesanywhere_deleteAttributeMappingCmd.Flags().String("specifiers", "", "A list of specifiers of a certificate field; for example, CN, OU, UID from a Subject.")
	rolesanywhere_deleteAttributeMappingCmd.MarkFlagRequired("certificate-field")
	rolesanywhere_deleteAttributeMappingCmd.MarkFlagRequired("profile-id")
	rolesanywhereCmd.AddCommand(rolesanywhere_deleteAttributeMappingCmd)
}
