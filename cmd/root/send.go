package root

import (
  "broadcast/pkg/send"

  "github.com/spf13/cobra"
)

var tags string

var sendCmd = &cobra.Command{
  Use:   "send",
  Aliases: []string{"s"},
  Short:  "Send a string",
  //Args:  cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    send.Broadcast(tags)
  },
}

func init() {
  sendCmd.Flags().StringVarP(&tags, "tags", "t", "", "Tags")
  rootCmd.AddCommand(sendCmd)
}
