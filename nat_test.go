func TestNetworkLoopbackNat(t *testing.T) {
	skip.If(t, testEnv.GitHubActions) // https://github.com/moby/moby/issues/41561
	skip.If(t, testEnv.OSType == "windows", "FIXME")
	skip.If(t, testEnv.IsRemoteDaemon)
