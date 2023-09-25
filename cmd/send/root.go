package send

import (
  "fmt"
  "os"

  "github.com/spf13/cobra"
  "github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
  Use:  "broadcast",
  Short: "broadcast - a simple CLI to transform and inspect strings",
  Long: `broadcast is a super fancy CLI

  Example: ...`,
  //Args:  cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
  },
}

func setDefaultCommandIfNonePresent() {
  cmd, _, err := rootCmd.Find(os.Args[1:])
  // default cmd if no cmd is given
  if err == nil && cmd.Use == rootCmd.Use && cmd.Flags().Parse(os.Args[1:]) != pflag.ErrHelp {
    args := append([]string{sendCmd.Use}, os.Args[1:]...)
    rootCmd.SetArgs(args)
  }
}

func Execute() {
  setDefaultCommandIfNonePresent()

  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
    os.Exit(1)
  }
}
