Windows PowerShell
Copyright (C) Microsoft Corporation. All rights reserved.

PS C:\1Go\SmallPrograms\gobug> go test ./... -v
# gobug
main.go:6:2: missing go.sum entry for module providing package github.com/wailsapp/wails/v2 (imported by gobug); to add:
        go get gobug
# gobug
main.go:7:2: missing go.sum entry for module providing package github.com/wailsapp/wails/v2/pkg/options (imported by gobug); to add:
        go get gobug
# gobug
main.go:8:2: missing go.sum entry for module providing package github.com/wailsapp/wails/v2/pkg/options/assetserver (imported by gobug); to add:
        go get gobug
FAIL    gobug [setup failed]
=== RUN   TestSaveAndLoad
--- PASS: TestSaveAndLoad (0.03s)
=== RUN   TestLoad_NoConfigYet
--- PASS: TestLoad_NoConfigYet (0.00s)
PASS
ok      gobug/config    1.731s
=== RUN   TestExplainWithAI_NoKey
--- PASS: TestExplainWithAI_NoKey (0.00s)
=== RUN   TestExtractContext_WithLineNumber
--- PASS: TestExtractContext_WithLineNumber (0.00s)
=== RUN   TestExtractContext_NoLineNumber
--- PASS: TestExtractContext_NoLineNumber (0.00s)
=== RUN   TestExplain_KnownPatterns
=== RUN   TestExplain_KnownPatterns/undefined_identifier
=== RUN   TestExplain_KnownPatterns/type_mismatch
=== RUN   TestExplain_KnownPatterns/unused_import
=== RUN   TestExplain_KnownPatterns/unused_variable
=== RUN   TestExplain_KnownPatterns/missing_return
=== RUN   TestExplain_KnownPatterns/nil_pointer_panic
=== RUN   TestExplain_KnownPatterns/index_out_of_range
=== RUN   TestExplain_KnownPatterns/unrecognized_error
--- PASS: TestExplain_KnownPatterns (0.00s)
    --- PASS: TestExplain_KnownPatterns/undefined_identifier (0.00s)
    --- PASS: TestExplain_KnownPatterns/type_mismatch (0.00s)
    --- PASS: TestExplain_KnownPatterns/unused_import (0.00s)
    --- PASS: TestExplain_KnownPatterns/unused_variable (0.00s)
    --- PASS: TestExplain_KnownPatterns/missing_return (0.00s)
    --- PASS: TestExplain_KnownPatterns/nil_pointer_panic (0.00s)
    --- PASS: TestExplain_KnownPatterns/index_out_of_range (0.00s)
    --- PASS: TestExplain_KnownPatterns/unrecognized_error (0.00s)
PASS
ok      gobug/engine    2.826s
FAIL
PS C:\1Go\SmallPrograms\gobug>