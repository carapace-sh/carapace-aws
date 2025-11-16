package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_getIdentityDkimAttributesCmd = &cobra.Command{
	Use:   "get-identity-dkim-attributes",
	Short: "Returns the current status of Easy DKIM signing for an entity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_getIdentityDkimAttributesCmd).Standalone()

	ses_getIdentityDkimAttributesCmd.Flags().String("identities", "", "A list of one or more verified identities - email addresses, domains, or both.")
	ses_getIdentityDkimAttributesCmd.MarkFlagRequired("identities")
	sesCmd.AddCommand(ses_getIdentityDkimAttributesCmd)
}
