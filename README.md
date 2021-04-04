# Using klog in a cobra application

Cobra uses pflags but the klog pkg forces the use of the standard go flags system.

This is a quick example how to configure klog to work with a cobra application.

Basic updates:
```diff
 import (
        "os"
 
+       goflags "flag"
+
        "github.com/spf13/cobra"
+       "k8s.io/klog/v2"
```
```diff
 func init() {
+       fs := goflags.NewFlagSet("", goflags.PanicOnError)
+       klog.InitFlags(fs)
+       rootCmd.Flags().AddGoFlagSet(fs)
+
```

Read the documentation (source code) for further details.
