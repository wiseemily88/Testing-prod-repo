func TestNetworkLoopbackNat(t *testing.T) {
	skip.If(t, testEnv.GitHubActions, "FIXME: Flaky test, see https://github.com//issues/9")
	skip.If(t, testEnv.OSType == "windows", "FIXME")
	skip.If(t, testEnv.IsRemoteDaemon)

	adding more here!
