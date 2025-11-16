package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisVideoArchivedMedia_listFragmentsCmd = &cobra.Command{
	Use:   "list-fragments",
	Short: "Returns a list of [Fragment]() objects from the specified stream and timestamp range within the archived data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisVideoArchivedMedia_listFragmentsCmd).Standalone()

	kinesisVideoArchivedMedia_listFragmentsCmd.Flags().String("fragment-selector", "", "Describes the timestamp range and timestamp origin for the range of fragments to return.")
	kinesisVideoArchivedMedia_listFragmentsCmd.Flags().String("max-results", "", "The total number of fragments to return.")
	kinesisVideoArchivedMedia_listFragmentsCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
	kinesisVideoArchivedMedia_listFragmentsCmd.Flags().String("stream-arn", "", "The Amazon Resource Name (ARN) of the stream from which to retrieve a fragment list.")
	kinesisVideoArchivedMedia_listFragmentsCmd.Flags().String("stream-name", "", "The name of the stream from which to retrieve a fragment list.")
	kinesisVideoArchivedMediaCmd.AddCommand(kinesisVideoArchivedMedia_listFragmentsCmd)
}
