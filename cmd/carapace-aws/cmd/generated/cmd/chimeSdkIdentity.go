package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentityCmd = &cobra.Command{
	Use:   "chime-sdk-identity",
	Short: "The Amazon Chime SDK Identity APIs in this section allow software developers to create and manage unique instances of their messaging applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentityCmd).Standalone()

	rootCmd.AddCommand(chimeSdkIdentityCmd)
}
