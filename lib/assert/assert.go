package assert

type TestIface interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fail()
	FailNow()
	Failed() bool
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Helper()
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Name() string
	Skip(args ...interface{})
	SkipNow()
	Skipf(format string, args ...interface{})
	Skipped() bool
}

func AssertEqual(t TestIface, a, b interface{}) {
	t.Helper()
	if a != b {
		t.Errorf("Not Equal. %d %d", a, b)
	}
}

func AssertTrue(t TestIface, a bool) {
	t.Helper()
	if !a {
		t.Errorf("Not True %t", a)
	}
}

func AssertFalse(t TestIface, a bool) {
	t.Helper()
	if a {
		t.Errorf("Not True %t", a)
	}
}

func AssertNil(t TestIface, a interface{}) {
	t.Helper()
	if a != nil {
		t.Error("Not Nil")
	}
}

func AssertNotNil(t TestIface, a interface{}) {
	t.Helper()
	if a == nil {
		t.Error("Is Nil")
	}
}
