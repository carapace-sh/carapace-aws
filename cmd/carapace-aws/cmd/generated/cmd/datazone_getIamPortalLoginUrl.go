package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getIamPortalLoginUrlCmd = &cobra.Command{
	Use:   "get-iam-portal-login-url",
	Short: "Gets the data portal URL for the specified Amazon DataZone domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getIamPortalLoginUrlCmd).Standalone()

	datazone_getIamPortalLoginUrlCmd.Flags().String("domain-identifier", "", "the ID of the Amazon DataZone domain the data portal of which you want to get.")
	datazone_getIamPortalLoginUrlCmd.MarkFlagRequired("domain-identifier")
	datazoneCmd.AddCommand(datazone_getIamPortalLoginUrlCmd)
}
