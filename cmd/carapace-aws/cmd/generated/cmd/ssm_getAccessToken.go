package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_getAccessTokenCmd = &cobra.Command{
	Use:   "get-access-token",
	Short: "Returns a credentials set to be used with just-in-time node access.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_getAccessTokenCmd).Standalone()

	ssm_getAccessTokenCmd.Flags().String("access-request-id", "", "The ID of a just-in-time node access request.")
	ssm_getAccessTokenCmd.MarkFlagRequired("access-request-id")
	ssmCmd.AddCommand(ssm_getAccessTokenCmd)
}
