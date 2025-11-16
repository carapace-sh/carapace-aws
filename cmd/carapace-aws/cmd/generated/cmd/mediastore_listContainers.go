package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediastore_listContainersCmd = &cobra.Command{
	Use:   "list-containers",
	Short: "Lists the properties of all containers in AWS Elemental MediaStore.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediastore_listContainersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediastore_listContainersCmd).Standalone()

		mediastore_listContainersCmd.Flags().String("max-results", "", "Enter the maximum number of containers in the response.")
		mediastore_listContainersCmd.Flags().String("next-token", "", "Only if you used `MaxResults` in the first command, enter the token (which was included in the previous response) to obtain the next set of containers.")
	})
	mediastoreCmd.AddCommand(mediastore_listContainersCmd)
}
