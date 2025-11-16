package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_getMembershipCmd = &cobra.Command{
	Use:   "get-membership",
	Short: "Retrieves a specified membership for an identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_getMembershipCmd).Standalone()

	cleanrooms_getMembershipCmd.Flags().String("membership-identifier", "", "The identifier for a membership resource.")
	cleanrooms_getMembershipCmd.MarkFlagRequired("membership-identifier")
	cleanroomsCmd.AddCommand(cleanrooms_getMembershipCmd)
}
