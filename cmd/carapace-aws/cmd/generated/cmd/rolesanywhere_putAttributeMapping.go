package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_putAttributeMappingCmd = &cobra.Command{
	Use:   "put-attribute-mapping",
	Short: "Put an entry in the attribute mapping rules that will be enforced by a given profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_putAttributeMappingCmd).Standalone()

	rolesanywhere_putAttributeMappingCmd.Flags().String("certificate-field", "", "Fields (x509Subject, x509Issuer and x509SAN) within X.509 certificates.")
	rolesanywhere_putAttributeMappingCmd.Flags().String("mapping-rules", "", "A list of mapping entries for every supported specifier or sub-field.")
	rolesanywhere_putAttributeMappingCmd.Flags().String("profile-id", "", "The unique identifier of the profile.")
	rolesanywhere_putAttributeMappingCmd.MarkFlagRequired("certificate-field")
	rolesanywhere_putAttributeMappingCmd.MarkFlagRequired("mapping-rules")
	rolesanywhere_putAttributeMappingCmd.MarkFlagRequired("profile-id")
	rolesanywhereCmd.AddCommand(rolesanywhere_putAttributeMappingCmd)
}
